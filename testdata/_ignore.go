package options

type Option func(options *Simple)

func OptionStringVal(option string) Option {
	return func(options *Simple) {
		options.StringVal = option
	}
}
func OptionIntVal(option int) Option {
	return func(options *Simple) {
		options.IntVal = option
	}
}
