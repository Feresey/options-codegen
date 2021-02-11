package options

type Embedded struct {
	Embed struct {
		some   int
		fields string
		in     struct{}
		lined  []uint8
	}
}
