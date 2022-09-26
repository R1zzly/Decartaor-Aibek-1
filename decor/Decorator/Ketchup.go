package Decorator

type Ketchup struct {
	fastFood FastFood
}

func (K *Ketchup) GetDescription() string {
	return K.fastFood.GetDescription() + ", ketchup"
}

func (K *Ketchup) GetCost() float64 {
	return K.fastFood.GetCost() + 0.2
}

func NewKetchup(food FastFood) *Ketchup {
	K := new(Ketchup)
	K.fastFood = food
	return K
}
