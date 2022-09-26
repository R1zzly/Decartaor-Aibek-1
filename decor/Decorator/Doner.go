package Decorator

import "fmt"

type Doner struct {
	description string
	cost        float64
	quantity    int
}

func (D *Doner) GetDescription() string {
	return fmt.Sprintf("%v x %v", D.description, D.quantity)
}

func (D *Doner) GetCost() float64 {
	return D.cost * float64(D.quantity)
}

func NewDoner(quantity int) *Doner {
	D := new(Doner)
	D.description = "Doner with chicken"
	D.cost = 3.5
	D.quantity = quantity
	return D
}
