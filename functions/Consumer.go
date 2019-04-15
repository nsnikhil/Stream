package functions

type Consumer struct {
	Accept func(interface{})
}

func NewConsumer(f func(interface{})) *Consumer {
	return &Consumer{Accept: f}
}
