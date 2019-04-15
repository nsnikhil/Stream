package stream

import (
	"Collections/functions"
)

type Stream struct {
	elements []interface{}
}

func Of(e interface{}) *Stream {
	return &Stream{elements: interfaceSlice(e)}
}

func Generate(seed interface{}, supplier *functions.Supplier) *Stream {
	return &Stream{}
}

func (s Stream) Filter(predicate *functions.Predicate) Stream {
	t := Stream{}
	for _, e := range s.elements {
		if predicate.Test(e) {
			t.elements = append(t.elements, e)
		}
	}
	return t
}

func (s Stream) Maps(function *functions.Function) Stream {
	t := Stream{}
	for _, e := range s.elements {
		t.elements = append(t.elements, function.Apply(e))
	}
	return t
}

func (s Stream) ForEach(consumer *functions.Consumer) {
	for _, e := range s.elements {
		consumer.Accept(e)
	}
}

func (s Stream) Peek(consumer *functions.Consumer) Stream {
	for _, e := range s.elements {
		consumer.Accept(e)
	}
	return s
}
