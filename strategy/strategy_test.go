package strategy

import "testing"

func Test(t *testing.T) {
	t.Run("strategy: ", SaleStrategy)
}

func SaleStrategy(t *testing.T) {
	var strategy Strategy
	strategy = NewSaleStrategy(SaleTypeA)
	strategy.Promotion()

	strategy = NewSaleStrategy(SaleTypeB)
	strategy.Promotion()

	strategy = NewSaleStrategy(SaleTypeC)
	strategy.Promotion()
}