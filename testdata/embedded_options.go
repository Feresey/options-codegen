// DO NOT EDIT!!!

package options

type Option func(options *Embedded)

func OptionEmbed(option struct {
	some   int
	fields string
	in     struct{}
	lined  []uint8
}) Option {
	return func(options *Embedded) {
		options.Embed = option
	}
}
