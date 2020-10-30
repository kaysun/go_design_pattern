package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("decorator: ", Decorator)
}

func Decorator(t *testing.T) {
	var myself, drinkCoffee, eatFriedChicken HappinessIndex
	myself = &MySelf{}
	fmt.Println(fmt.Sprintf("我的幸福指数是：%d", myself.GetHappinessIndex()))

	drinkCoffee = &DrinkCoffee{HappinessIndex:MySelf{}}
	fmt.Println(fmt.Sprintf("喝了咖啡后，我的幸福指数是：%d", drinkCoffee.GetHappinessIndex()))

	eatFriedChicken = &EatFriedChicken{HappinessIndex:drinkCoffee}
	fmt.Println(fmt.Sprintf("吃了炸鸡，喝了咖啡后，我的幸福指数是：%d", eatFriedChicken.GetHappinessIndex()))
}