// package abstract_factory 抽象工厂模式
package abstract_factory

import "fmt"

// FruitFacotry 水果工厂接口
type FruitFacotry interface {
	// CreateFruit 生产水果
	CreateFruit() Fruit
}

//AppleFactory 苹果工厂，实现FruitFacotry接口
type AppleFactory struct{}

//BananaFactory 香蕉工厂，实现FruitFacotry接口
type BananaFactory struct{}

//OrangeFactory 橘子工厂，实现FruitFacotry接口
type OrangeFactory struct{}

// CreateFruit 苹果工厂生产苹果
func (appleFactory AppleFactory) CreateFruit() Fruit {
	return &Apple{}
}

// CreateFruit 香蕉工厂生产香蕉
func (bananaFactory BananaFactory) CreateFruit() Fruit {
	return &Banana{}
}

// CreateFruit 橘子工厂生产橘子
func (orangeFactory OrangeFactory) CreateFruit() Fruit {
	return &Orange{}
}

// Fruit 水果接口
type Fruit interface {
	// Eat 吃水果
	Eat()
}

// Apple 苹果，实现Fruit接口
type Apple struct{}

// Banana 香蕉，实现Fruit接口
type Banana struct{}

// Orange 橘子，实现Fruit接口
type Orange struct{}

// Eat 吃苹果
func (apple Apple) Eat() {
	fmt.Println("吃苹果")
}

// Eat 吃香蕉
func (banana Banana) Eat() {
	fmt.Println("吃香蕉")
}

// Eat 吃橘子
func (orange Orange) Eat() {
	fmt.Println("吃橘子")
}
