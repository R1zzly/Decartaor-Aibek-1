package Decorator

type DonerMeat struct {
	Doner
}

func NewDonerMeat(quantity int) *DonerMeat {
	D := new(DonerMeat)
	D.description = "Doner with meat"
	D.cost = 4
	D.quantity = quantity
	return D
}
