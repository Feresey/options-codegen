// DO NOT EDIT!!!

package options

type Option func(options *BaseDecls)

func OptionUnexportedEmptyVal(option struct{}) Option {
	return func(options *BaseDecls) {
		options.unexportedEmptyVal = option
	}
}
func OptionChanVal(option chan struct{}) Option {
	return func(options *BaseDecls) {
		options.ChanVal = option
	}
}
func OptionMapVal(option map[string]interface{}) Option {
	return func(options *BaseDecls) {
		options.MapVal = option
	}
}
func OptionSliceVal(option []int) Option {
	return func(options *BaseDecls) {
		options.SliceVal = option
	}
}
func OptionUnnamedFunctionVal(option func()) Option {
	return func(options *BaseDecls) {
		options.UnnamedFunctionVal = option
	}
}
func OptionAnyIface(option interface{}) Option {
	return func(options *BaseDecls) {
		options.AnyIface = option
	}
}
func OptionStaredVal(option *[]*[]*[]*[]********map[*int]*int) Option {
	return func(options *BaseDecls) {
		options.StaredVal = option
	}
}
