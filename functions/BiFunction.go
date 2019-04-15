package functions

type BiFunction struct {
	Apply func(interface{}, interface{}) interface{}
}

func NewBiFunction(f func(interface{}, interface{}) interface{}) *BiFunction {
	return &BiFunction{Apply: f}
}
