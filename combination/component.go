//package combination 组合模式
package combination

// UIComponent UI组件接口，对于任何UI控件都适用。
type UIComponent interface {
	// PrintUIComponent 打印UI组件
	PrintUIComponent()
	// GetUIControlName 获取控件名字
	GetUIControlName() string
	// GetConcreteUIControlName 获取控件具体名字
	GetConcreteUIControlName() string
}

// UIComponentAddtion UI组件附加接口，使用接口隔离原则，保证不需要实现接口声明方法的结构体，没有额外负担。仅对容器类型对UI控件适用。
type UIComponentAddtion interface {
	// AddUIComponent 添加UI组件
	AddUIComponent(component UIComponent)
	// AddUIComponents 添加UI组件列表
	AddUIComponents(components []UIComponent)
	// GetUIComponentList 获取UI组件列表
	GetUIComponentList() []UIComponent
}

// UIAttr UI属性
type UIAttr struct {
	// UI 名字
	Name string
}

//client 打印client
var client = &PrintClient{}
