package functions

type BiConsumer struct {
	Accept func(interface{}, interface{})
}

func NewBiConsumer(f func(interface{}, interface{})) *BiConsumer {
	return &BiConsumer{Accept: f}
}
