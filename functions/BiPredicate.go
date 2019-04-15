package functions

type BiPredicate struct {
	Test func(interface{}, interface{}) bool
}

func NewBiPredicate(f func(interface{}, interface{}) bool) *BiPredicate {
	return &BiPredicate{Test: f}
}
