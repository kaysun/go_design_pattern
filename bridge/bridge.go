// Package bridge 桥接模式，以订购咖啡为例，两个纬度，分别为大中小杯，加糖加奶
package bridge

import "fmt"

// ICoffee 咖啡接口
type ICoffee interface {
	// OrderCoffee 订购咖啡
	OrderCoffee()
}

// LargeCoffee 大杯咖啡，实现ICoffee接口
type LargeCoffee struct {
	ICoffeeAddtion
}

// MediumCoffee 中杯咖啡，实现ICoffee接口
type MediumCoffee struct {
	ICoffeeAddtion
}

// SmallCoffee 小杯咖啡，实现ICoffee接口
type SmallCoffee struct {
	ICoffeeAddtion
}

// OrderCoffee 订购大杯咖啡
func (lc LargeCoffee) OrderCoffee() {
	fmt.Println("订购了大杯咖啡")
	lc.AddSomething()
}

// OrderCoffee 订购中杯咖啡
func (mc MediumCoffee) OrderCoffee() {
	fmt.Println("订购了中杯咖啡")
	mc.AddSomething()
}

// OrderCoffee 订购小杯咖啡
func (sc SmallCoffee) OrderCoffee() {
	fmt.Println("订购了小杯咖啡")
	sc.AddSomething()
}

// CoffeeCupType 咖啡容量类型
type CoffeeCupType uint8

const (
	// CoffeeCupTypeLarge 大杯咖啡
	CoffeeCupTypeLarge = iota
	// CoffeeCupTypeMedium 中杯咖啡
	CoffeeCupTypeMedium = iota
	// CoffeeCupTypeSmall 小杯咖啡
	CoffeeCupTypeSmall = iota
)

// CoffeeFuncMap 全局可导出变量，咖啡类型与创建咖啡对象的map，用于减小圈复杂度
var CoffeeFuncMap = map[CoffeeCupType]func(coffeeAddtion ICoffeeAddtion) ICoffee{
	CoffeeCupTypeLarge:  NewLargeCoffee,
	CoffeeCupTypeMedium: NewMediumCoffee,
	CoffeeCupTypeSmall:  NewSmallCoffee,
}

// NewCoffee 创建咖啡接口对象的简单工厂，根据咖啡容量类型，获取创建接口对象的func
func NewCoffee(cupType CoffeeCupType, coffeeAddtion ICoffeeAddtion) ICoffee {
	if handler, ok := CoffeeFuncMap[cupType]; ok {
		return handler(coffeeAddtion)
	}
	return nil
}

// NewLargeCoffee 创建大杯咖啡对象
func NewLargeCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &LargeCoffee{coffeeAddtion}
}

// NewMediumCoffee 创建中杯咖啡对象
func NewMediumCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &MediumCoffee{coffeeAddtion}
}

// NewSmallCoffee 创建小杯咖啡对象
func NewSmallCoffee(coffeeAddtion ICoffeeAddtion) ICoffee {
	return &SmallCoffee{coffeeAddtion}
}

// ICoffeeAddtion 咖啡额外添加接口
type ICoffeeAddtion interface {
	//AddSomething 添加某种食物
	AddSomething()
}

// Milk 加奶，实现ICoffeeAddtion接口
type Milk struct{}

// Sugar 加糖，实现ICoffeeAddtion接口
type Sugar struct{}

// AddSomething Milk实现加奶
func (milk Milk) AddSomething() {
	fmt.Println("加奶")
}

// AddSomething Sugar实现加糖
func (sugar Sugar) AddSomething() {
	fmt.Println("加糖")
}

// CoffeeAddtionType 咖啡额外添加类型
type CoffeeAddtionType uint8

const (
	// CoffeeAddtionTypeMilk 咖啡额外添加牛奶
	CoffeeAddtionTypeMilk = iota
	// CoffeeAddtionTypeSugar 咖啡额外添加糖
	CoffeeAddtionTypeSugar = iota
)

// CoffeeAddtionFuncMap 全局可导出变量，咖啡额外添加类型与创建咖啡额外添加对象的map，用于减小圈复杂度
var CoffeeAddtionFuncMap = map[CoffeeAddtionType]func() ICoffeeAddtion{
	CoffeeAddtionTypeMilk:  NewCoffeeAddtionMilk,
	CoffeeAddtionTypeSugar: NewCoffeeAddtionSugar,
}

// NewCoffeeAddtion 创建咖啡额外添加接口对象的简单工厂，根据咖啡额外添加类型，获取创建接口对象的func
func NewCoffeeAddtion(addtionType CoffeeAddtionType) ICoffeeAddtion {
	if handler, ok := CoffeeAddtionFuncMap[addtionType]; ok {
		return handler()
	}
	return nil
}

// NewCoffeeAddtionMilk 创建咖啡额外加奶
func NewCoffeeAddtionMilk() ICoffeeAddtion {
	return &Milk{}
}

// NewCoffeeAddtionSugar 创建咖啡额外加糖
func NewCoffeeAddtionSugar() ICoffeeAddtion {
	return &Sugar{}
}
