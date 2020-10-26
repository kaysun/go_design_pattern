// package strategy 策略模式。以商场促销方案为例。
package strategy

import "fmt"

// Strategy 策略接口
type Strategy interface {
	// Promotion 促销
	Promotion()
}

// ConcreteStrategyA 促销策略A
type ConcreteStrategyA struct {

}

// ConcreteStrategyB 促销策略B
type ConcreteStrategyB struct {

}

// ConcreteStrategyC 促销策略C
type ConcreteStrategyC struct {

}

func (strategy ConcreteStrategyA)Promotion()  {
	fmt.Println("618促销")
}

func (strategy ConcreteStrategyB)Promotion()  {
	fmt.Println("99促销")
}

func (strategy ConcreteStrategyC)Promotion()  {
	fmt.Println("双十一促销")
}
