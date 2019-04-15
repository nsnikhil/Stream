package functions

type Predicate struct {
	Test func(interface{}) bool
}

func NewPredicate(f func(i interface{}) bool) *Predicate {
	return &Predicate{Test: f}
}


