// DO NOT EDIT!!!

package options

import (
	"context"
	"fmt"
	jen "github.com/dave/jennifer/jen"
)

type Option func(options *ExtendedDecls)

func OptionAliasedVal(option string) Option {
	return func(options *ExtendedDecls) {
		options.AliasedVal = option
	}
}
func OptionNamedVal(option *NameType) Option {
	return func(options *ExtendedDecls) {
		options.NamedVal = option
	}
}
func OptionSturctVal(option *SimpleStruct) Option {
	return func(options *ExtendedDecls) {
		options.SturctVal = option
	}
}
func OptionFunctionVal(option context.CancelFunc) Option {
	return func(options *ExtendedDecls) {
		options.FunctionVal = option
	}
}
func OptionExternIface(option fmt.Stringer) Option {
	return func(options *ExtendedDecls) {
		options.ExternIface = option
	}
}
func OptionExternVal(option *jen.File) Option {
	return func(options *ExtendedDecls) {
		options.ExternVal = option
	}
}
