package options

import (
	"context"
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

type SimpleStruct struct {
	Field int
	Array []int
}

type NameType strings.Replacer

type AliasType = string

type ExtendedDecls struct {
	AliasedVal  AliasType
	NamedVal    *NameType
	SturctVal   *SimpleStruct
	FunctionVal context.CancelFunc
	ExternIface fmt.Stringer
	ExternVal   *jen.File
}
