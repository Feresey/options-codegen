package options

type BaseDecls struct {
	unexportedEmptyVal struct{}
	ChanVal            chan struct{}
	MapVal             map[string]interface{}
	SliceVal           []int
	UnnamedFunctionVal func()
	AnyIface           interface{}
	StaredVal          *[]*[]*[]*[]********map[*int]*int
}
