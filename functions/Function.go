package functions

type Function struct {
	Apply func(interface{}) interface{}
}

func NewFunction(f func(interface{}) interface{}) *Function {
	return &Function{Apply: f}
}
