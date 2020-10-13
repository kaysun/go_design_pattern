package simple_factory

import "testing"

func Test(t *testing.T) {
	t.Run("simple_factory: ", ProduceFruitAndEat)
}

func ProduceFruitAndEat(t *testing.T) {
	var apple, banana, orange Fruit
	apple = ProduceFruit(FruitTypeApple)
	banana = ProduceFruit(FruitTypeBanana)
	orange = ProduceFruit(FruitTypeOrange)

	apple.Peeling()
	apple.Eat()

	banana.Peeling()
	banana.Eat()

	orange.Peeling()
	orange.Eat()
}