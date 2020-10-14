package factory

import "testing"

func Test(t *testing.T) {
	t.Run("factory: ", ProduceFruitAndEat)
}

func ProduceFruitAndEat(t *testing.T) {
	appleFactory := &AppleFactory{}
	bananaFactory := &BananaFactory{}
	orangeFactory := &OrangeFactory{}

	var apple, banana, orange Fruit

	apple = appleFactory.CreateFruit()
	apple.Eat()

	banana = bananaFactory.CreateFruit()
	banana.Eat()

	orange = orangeFactory.CreateFruit()
	orange.Eat()
}