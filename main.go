package main

import (
	"fmt"
	"go/types"

	"github.com/spf13/pflag"
	"golang.org/x/tools/go/packages"
)

type config struct {
	StructName string
	Input      string
	Output     string
}

type generator struct {
	config

	pkg *packages.Package
}

func main() {
	var g generator
	pflag.StringVarP(&g.StructName, "struct", "s", "", "struct name")
	pflag.StringVarP(&g.Input, "input", "i", "", "input package or file, may be ${GOFILE}")
	pflag.StringVarP(&g.Output, "output", "o", "", "output file name, default is {struct_name}_options.go")
	pflag.Parse()

	if err := g.Run(); err != nil {
		panic(err)
	}
}

func (g *generator) Run() error {
	g.Init()
	if err := g.Parse(); err != nil {
		return err
	}

	// if err := g.GetFields()

	return nil
}

func (g *generator) Init() {
	if g.config.Output == "" {
		g.config.Output = g.config.StructName + "_options.go"
	}
}

func (g *generator) Parse() error {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedTypes,
	}, g.Input)
	if err != nil {
		return fmt.Errorf("parse packages: %w", err)
	}
	if len(pkgs) != 0 {
		return fmt.Errorf("number of packages must be one, got: %d", len(pkgs))
	}
	g.pkg = pkgs[0]

	return nil
}

type Field struct {
	Name string
	// Flags []string
	Type *types.Type
}

func (g *generator) GetFields() error {
	object := g.pkg.Types.Scope().Lookup(g.StructName)
	if object == nil {
		return fmt.Errorf("object not foun in scope: %q", g.StructName)
	}

	desc, ok := object.Type().(*types.Struct)
	if !ok {
		return fmt.Errorf("%q is not struct type: [%s]", g.StructName, object.Type().String())
	}

	for i := 0; i < desc.NumFields(); i++ {
		field := desc.Field(i)
		_ = field
	}

	return nil
}
