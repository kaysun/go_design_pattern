// package decorator 装饰者模式，在运行期动态地给对象添加额外的指责，比子类更灵活。第一个功能：用于增强功能的装饰模式。
package decorator

// HappinessIndex 幸福指数接口
type HappinessIndex interface {
	// GetHappinessIndex 获取幸福指数
	GetHappinessIndex() int
}

// MySelf 我自己，实现HappinessIndex接口
type MySelf struct {}

// GetHappinessIndex 获取幸福指数，实现HappinessIndex接口
func (myself MySelf) GetHappinessIndex() int {
	return 100
}

// DrinkCoffee 喝咖啡
type DrinkCoffee struct {
	// HappinessIndex 幸福指数对象
	HappinessIndex HappinessIndex
}

// GetHappinessIndex 获取幸福指数，实现HappinessIndex接口
func (coffee DrinkCoffee) GetHappinessIndex() int {
	return coffee.HappinessIndex.GetHappinessIndex() + 20
}

// EatFriedChicken 吃炸鸡，实现HappinessIndex接口
type EatFriedChicken struct {
	// HappinessIndex 幸福指数对象
	HappinessIndex HappinessIndex
}

// GetHappinessIndex 获取幸福指数，实现HappinessIndex接口
func (friedChicken EatFriedChicken) GetHappinessIndex() int {
	return friedChicken.HappinessIndex.GetHappinessIndex() + 50
}