package Decorator

type Mayo struct {
	fastFood FastFood
}

func (K *Mayo) GetDescription() string {
	return K.fastFood.GetDescription() + ", mayo"
}

func (K *Mayo) GetCost() float64 {
	return K.fastFood.GetCost() + 0.2
}

func NewMayo(food FastFood) *Mayo {
	K := new(Mayo)
	K.fastFood = food
	return K
}
