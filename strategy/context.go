package strategy

import "context"

// Context
type Context struct {
	// context.Context 嵌套上下文
	context.Context
	//Strategy 策略对象
	Strategy Strategy
}

// Sale 促销
func Sale(saleType SaleType) {
	strategy := NewSaleStrategy(saleType)
	strategy.Promotion()
}

// SaleType 促销类型
type SaleType uint8

const (
	// SaleTypeA 促销类型A
	SaleTypeA SaleType = iota
	// SaleTypeB 促销类型B
	SaleTypeB
	// SaleTypeC 促销类型C
	SaleTypeC
)

// SaleTypeFuncMap 全局可导出变量，咖啡额外添加类型与创建咖啡额外添加对象的map，用于减小圈复杂度
var SaleTypeFuncMap = map[SaleType]func() Strategy {
	SaleTypeA: NewStrategyA,
	SaleTypeB: NewStrategyB,
	SaleTypeC: NewStrategyC,
}

// NewSaleStrategy 创建促销类型接口对象的简单工厂，根据促销类型类型，获取创建接口对象的func
func NewSaleStrategy(saleType SaleType) Strategy {
	if handler, ok := SaleTypeFuncMap[saleType]; ok {
		return handler()
	}
	return nil
}

// NewStrategyA 创建促销A
func NewStrategyA() Strategy {
	return &ConcreteStrategyA{}
}

// NewStrategyB 创建促销B
func NewStrategyB() Strategy {
	return &ConcreteStrategyB{}
}

// NewStrategyC 创建促销C
func NewStrategyC() Strategy {
	return &ConcreteStrategyC{}
}