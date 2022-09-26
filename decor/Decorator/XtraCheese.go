package Decorator

type Cheese struct {
	fastFood FastFood
}

func (K *Cheese) GetDescription() string {
	return K.fastFood.GetDescription() + ", extra cheese"
}

func (K *Cheese) GetCost() float64 {
	return K.fastFood.GetCost() + 0.5
}

func NewCheese(food FastFood) *Cheese {
	K := new(Cheese)
	K.fastFood = food
	return K
}
