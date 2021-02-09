// DO NOT EDIT!!!

package options

type Option func(options *Simple)

func OptionIntVal(option int) Option {
	return func(options *Simple) {
		options.IntVal = option
	}
}
