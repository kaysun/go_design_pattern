package decorator

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("decorator_strengthen: ", DecoratorStrengthen)
	t.Run("decorator_add: ", DecoratorAdd)
}

func DecoratorStrengthen(t *testing.T) {
	var myself, drinkCoffee, eatFriedChicken HappinessIndex
	myself = &MySelf{}
	fmt.Println(fmt.Sprintf("我的幸福指数是：%d", myself.GetHappinessIndex()))

	drinkCoffee = &DrinkCoffee{HappinessIndex: MySelf{}}
	fmt.Println(fmt.Sprintf("喝了咖啡后，我的幸福指数是：%d", drinkCoffee.GetHappinessIndex()))

	eatFriedChicken = &EatFriedChicken{HappinessIndex: drinkCoffee}
	fmt.Println(fmt.Sprintf("吃了炸鸡，喝了咖啡后，我的幸福指数是：%d", eatFriedChicken.GetHappinessIndex()))
	fmt.Println("============")
}

func DecoratorAdd(t *testing.T) {
	var book Booker
	book = &Book{}
	book.Reading()
	fmt.Println("============")

	var notesTake NotesTaker
	notesTake = &ConcreteNotesTake{Booker: book}
	notesTake.Reading()
	notesTake.TakeNotes()
	fmt.Println("============")

	var Underline Underliner
	Underline = &ConcreteUnderline{Booker: book}
	Underline.Reading()
	Underline.Underline()
}
