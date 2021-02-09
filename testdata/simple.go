//nolint // this is test
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

type Simple struct {
	AliasedVal AliasType

	//options:ignore
	StringVal string
	IntVal    int

	SturctVal          *SimpleStruct
	ChanVal            chan struct{}
	MapVal             map[string]interface{}
	SliceVal           []int
	FunctionVal        context.CancelFunc
	UnnamedFunctionVal func()
	NamedVal           *NameType
	StaredVal          *[]*[]*[]*[]********map[*int]*int

	ExternVal *jen.File

	AnyIface    interface{}
	ExternIface fmt.Stringer

	unexportedEmptyVal struct{}
}
