package options

type SimpleStruct struct {
	Field int
	Array []int
}

type Simple struct {
	//options:ignore
	StringVal string
	SturctVal *SimpleStruct
	IntVal    int
	AnyVal    interface{}

	unexportedEmptyVal struct{} //nolint:structcheck,unused // this is test
}
