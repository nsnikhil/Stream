package stream

import (
	"Collections/functions"
	"sync"
)

type LazyStream struct {
	elements []interface{}
	wg       sync.WaitGroup
}

func OfLazy(e interface{}) *LazyStream {
	return &LazyStream{elements: interfaceSlice(e)}
}

func (ls *LazyStream) LazyFilter(predicate *functions.Predicate) *LazyStream {
	ls.wg.Add(1)
	v := <-ls.lazyFilter(predicate)
	return &v
}

func (ls *LazyStream) lazyFilter(predicate *functions.Predicate) chan LazyStream {
	c := make(chan LazyStream)
	t := LazyStream{}
	go func() {
		for _, e := range ls.elements {
			if predicate.Test(e) {
				t.elements = append(t.elements, e)
			}
		}
		c <- t
		ls.wg.Done()
	}()
	return c
}

func (ls *LazyStream) LazyMap(function *functions.Function) *LazyStream {
	ls.wg.Add(1)
	v := <-ls.lazyMap(function)
	return &v
}

func (ls *LazyStream) lazyMap(function *functions.Function) chan LazyStream {
	c := make(chan LazyStream)
	t := LazyStream{}
	go func() {
		for _, e := range ls.elements {
			t.elements = append(t.elements, function.Apply(e))
		}
		c <- t
		ls.wg.Done()
	}()
	return c
}

func (ls *LazyStream) LazyForEach(consumer *functions.Consumer) {
	ls.wg.Wait()
	for _, e := range ls.elements {
		consumer.Accept(e)
	}
}
