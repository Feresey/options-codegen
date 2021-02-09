package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"golang.org/x/tools/go/packages"
)

type config struct {
	StructName       string
	Input            string
	Output           string
	OptionTypename   string
	OptionFuncPrefix string

	ignoreOutput bool
}

type generator struct {
	config

	fields    []Field
	fieldsMap map[string]*ast.Field

	pkg *packages.Package
}

func main() {
	var g generator
	pflag.StringVarP(&g.StructName, "struct", "s", "", "struct name")
	pflag.StringVarP(&g.Input, "input", "i", "", "input package")
	pflag.StringVar(&g.OptionTypename, "typename", "Option", "name of result option type")
	pflag.StringVar(&g.OptionFuncPrefix, "prefix", "Option", "prefix of the option functions")
	pflag.StringVarP(&g.Output, "output", "o", "", "output file name, default is {struct_name}_options.go")
	pflag.Parse()

	if err := g.Run(); err != nil {
		panic(err)
	}
}

func (g *generator) Run() error {
	g.Init()
	if err := g.Parse(); err != nil {
		return fmt.Errorf("parse: %w", err)
	}

	if err := g.GetFields(); err != nil {
		return fmt.Errorf("get fields: %w", err)
	}

	file := g.Generate()

	out := g.Output
	if !filepath.IsAbs(g.Output) {
		out = filepath.Join(g.Input, g.Output)
	}

	if g.config.ignoreOutput {
		if err := file.Save(out); err != nil {
			return fmt.Errorf("generate options: %w", err)
		}
	} else {
		return file.Render(ioutil.Discard)
	}

	return nil
}

func (g *generator) Init() {
	if g.config.Output == "" {
		g.config.Output = strcase.ToSnake(g.config.StructName) + "_options.go"
	}
}

func (g *generator) Parse() error {
	absPath, err := filepath.Abs(g.Input)
	if err != nil {
		return fmt.Errorf("get full path of %q: %w", g.Input, err)
	}
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes | packages.NeedImports | packages.NeedName | packages.NeedSyntax,
	}, absPath)
	if err != nil {
		return fmt.Errorf("parse packages: %w", err)
	}
	if len(pkgs) != 1 {
		return fmt.Errorf("number of packages must be one, got: %d", len(pkgs))
	}
	g.pkg = pkgs[0]

	return nil
}

type Field struct {
	Name string
	// Flags []string
	Type types.Type
}

func (g *generator) GetFields() error {
	object := g.pkg.Types.Scope().Lookup(g.StructName)
	if object == nil {
		return fmt.Errorf("object not found in scope: %q", g.StructName)
	}

	desc, ok := object.Type().Underlying().(*types.Struct)
	if !ok {
		return fmt.Errorf("%q is not struct type: [%s]", g.StructName, object.Type().String())
	}

	g.getFieldsMap()
FIELDS:
	for i := 0; i < desc.NumFields(); i++ {
		field := desc.Field(i)
		astField := g.fieldsMap[field.Name()]
		if astField != nil {
			if doc := astField.Doc; doc != nil {
				for _, line := range doc.List {
					tag := g.parseTag(line.Text)
					if tag == tagSkip {
						continue FIELDS
					}
				}
			}
		}

		option := Field{
			Name: field.Name(),
			Type: field.Type(),
		}

		g.fields = append(g.fields, option)
	}

	return nil
}

func (g *generator) getFieldsMap() {
	g.fieldsMap = make(map[string]*ast.Field)
	for _, file := range g.pkg.Syntax {
		obj := file.Scope.Lookup(g.StructName)
		if obj == nil {
			continue
		}
		typ := obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType)
		list := typ.Fields.List
		for _, field := range list {
			// TODO names???
			g.fieldsMap[field.Names[0].String()] = field
		}
		return //nolint
	}
}

//go:generate go run github.com/alvaroloes/enumer -type tag -trimprefix tag
type tag int

const (
	tagNothig tag = iota
	tagSkip
)

const tagPrefix = "//options:"

func (g *generator) parseTag(line string) tag {
	if strings.HasPrefix(line, tagPrefix) {
		res, _ := tagString(strings.TrimPrefix(line, tagPrefix))
		return res
	}
	return tagNothig
}

func (g *generator) Generate() *jen.File {
	f := jen.NewFile(g.pkg.Name)

	structID := jen.Id("options")

	opFuncType := jen.Func().Params(
		structID.Clone().Op("*").Qual("", g.StructName),
	)

	f.Type().Id(g.OptionTypename).Add(opFuncType)

	optionType := jen.Id(g.OptionTypename)

	for _, field := range g.fields {
		typ := field.Type
		path, typename, stared := g.getType(typ)

		optionID := jen.Id("option")
		option := optionID.Clone()
		for i := 0; i < stared; i++ {
			option = option.Op("*")
		}
		option = option.Qual(path, typename)

		f.Func().
			Id(g.OptionFuncPrefix + strcase.ToCamel(field.Name)).
			Params(option).List(optionType).
			Block(
				jen.Return(opFuncType.Clone().
					Block(
						structID.Clone().
							Dot(field.Name).
							Op("=").
							Add(optionID.Clone()),
					)),
			)
	}

	return f
}

func (g *generator) getType(typ types.Type) (path, typename string, stared int) {
	switch typ := typ.(type) {
	case
		*types.Signature,
		*types.Array,
		*types.Basic,
		*types.Chan,
		*types.Tuple,
		*types.Slice,
		*types.Struct,
		*types.Map,
		*types.Interface:
		typename = typ.String()
	case *types.Named:
		named := typ.Obj()
		if named.IsAlias() {
			spew.Dump(named)
		}
		if named.Pkg() != g.pkg.Types {
			path = named.Pkg().Path()
		}
		typename = named.Name()
	case *types.Pointer:
		path, typename, stared = g.getType(typ.Elem())
		stared++
	default:
		panic("unsupported type")
	}
	return
}
