package functions

type Supplier struct {
	Get func() interface{}
}

func NewSupplier(f func() interface{}) *Supplier {
	return &Supplier{Get: f}
}
