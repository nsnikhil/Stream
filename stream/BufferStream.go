package stream

import (
	"Collections/functions"
)

type BufferStream struct {
	elements   []interface{}
	operations []interface{}
}

func OfBuffer(e interface{}) *BufferStream {
	return &BufferStream{elements: interfaceSlice(e)}
}

func (bfs *BufferStream) BufferFilter(predicate *functions.Predicate) *BufferStream {
	bfs.operations = append(bfs.operations, predicate)
	return bfs
}

func (bfs *BufferStream) BufferMap(function *functions.Function) *BufferStream {
	bfs.operations = append(bfs.operations, function)
	return bfs
}

func (bfs *BufferStream) BufferForEach(consumer *functions.Consumer) {
	bfs.operations = append(bfs.operations, consumer)
	bfs.runTermination()
}

func (bfs *BufferStream) runTermination() {
	for _, o := range bfs.operations {
		switch o.(type) {
		case *functions.Predicate:
			bfs.runPredicate(o.(*functions.Predicate))
		case *functions.Function:
			bfs.runFunction(o.(*functions.Function))
		case *functions.Consumer:
			bfs.runConsumer(o.(*functions.Consumer))
		}
	}
}

func (bfs *BufferStream) runPredicate(predicate *functions.Predicate) {
	var el []interface{}
	for _, e := range bfs.elements {
		if predicate.Test(e) {
			el = append(el, e)
		}
	}
	bfs.elements = el
}

func (bfs *BufferStream) runFunction(function *functions.Function) {
	var el []interface{}
	for _, e := range bfs.elements {
		el = append(el, function.Apply(e))
	}
	bfs.elements = el
}

func (bfs *BufferStream) runConsumer(consumer *functions.Consumer) {
	for _, e := range bfs.elements {
		consumer.Accept(e)
	}
}
