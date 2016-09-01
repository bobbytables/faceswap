package dummy

type CustomType struct{}

// Method names on interfaces are sorted by the types package in Go. So we're
// just prepending a letter to keep the same order.
type FakeInterface interface {
	A_Hello(world string)
	B_AnError() error
	C_MultiReturn() (string, error)
	D_MultiReturnCustomType() (*CustomType, error)
	E_CustomReturn(ce *CustomType) *CustomType
}
