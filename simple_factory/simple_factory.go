// package simple_factory 简单工厂模式
package simple_factory

import "fmt"

// Fruit 水果接口
type Fruit interface {
	// Peeling 削果皮
	Peeling()

	// Eat 吃水果
	Eat()
}

// Apple 苹果，实现Fruit接口
type Apple struct {}

// Banana 香蕉，实现Fruit接口
type Banana struct {}

// Orange 橘子，实现Fruit接口
type Orange struct {}

func (apple Apple) Peeling() {
	fmt.Println("削苹果果皮")
}

func (apple Apple) Eat() {
	fmt.Println("吃苹果")
}

func (banana Banana) Peeling() {
	fmt.Println("剥开香蕉果皮")
}

func (banana Banana) Eat() {
	fmt.Println("吃香蕉")
}

func (orange Orange) Peeling() {
	fmt.Println("剥开橘子果皮")
}

func (orange Orange) Eat() {
	fmt.Println("吃橘子")
}

// FruitType 水果类型
type FruitType uint8

const (
	// FruitTypeApple 苹果
	FruitTypeApple = iota
	// FruitTypeBanana 香蕉
	FruitTypeBanana = iota
	// FruitTypeOrange 橙子
	FruitTypeOrange = iota
)

// FruitFuncMap 全局可导出变量，水果类型与创建水果对象的map，用于减小圈复杂度
var FruitFuncMap = map[FruitType]func() Fruit {
	FruitTypeApple: produceApple,
	FruitTypeBanana: produceBanana,
	FruitTypeOrange: produceOrange,
}

// ProduceFruit 生产水果的简单工厂，根据水果类型，获取创建接口对象的func
func ProduceFruit(fruitType FruitType) Fruit {
	if handler, ok := FruitFuncMap[fruitType]; ok {
		return handler()
	}
	return nil
}

// produceApple 生产苹果
func produceApple() Fruit {
	return &Apple{}
}

// produceApple 生产香蕉
func produceBanana() Fruit {
	return &Banana{}
}

// produceApple 生产橘子
func produceOrange() Fruit {
	return &Orange{}
}