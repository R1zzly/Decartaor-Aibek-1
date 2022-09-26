package main

import (
	"decor/Decorator"
	"fmt"
)

func main() {
	var ff Decorator.FastFood = Decorator.NewDoner(1)
	ff = Decorator.NewKetchup(ff)
	ff = Decorator.NewMayo(ff)
	ff = Decorator.NewCheese(ff)
	fmt.Printf("%v is %v $ \n", ff.GetDescription(), ff.GetCost())
	var ffm Decorator.FastFood = Decorator.NewDonerMeat(2)
	ffm = Decorator.NewKetchup(ffm)
	ffm = Decorator.NewCheese(ffm)
	fmt.Printf("%v is %v $ \n", ffm.GetDescription(), ffm.GetCost())
}
