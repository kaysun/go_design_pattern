package bridge

import "testing"

func Test(t *testing.T) {
	t.Run("large-milk: ", OrderLargeMilkCoffee)
	t.Run("large-sugar: ", OrderLargeSugarCoffee)
	t.Run("medium-milk: ", OrderMediumMilkCoffee)
	t.Run("smarll-sugar: ", OrderSmallSugarCoffee)
}

func OrderLargeMilkCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeMilk)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeLarge, coffeeAddtion)
	coffee.OrderCoffee()
}

func OrderLargeSugarCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeSugar)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeLarge, coffeeAddtion)
	coffee.OrderCoffee()
}

func OrderMediumMilkCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeMilk)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeMedium, coffeeAddtion)
	coffee.OrderCoffee()
}

func OrderSmallSugarCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeSugar)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeSmall, coffeeAddtion)
	coffee.OrderCoffee()
}
