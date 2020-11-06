package combination

// WinForm 窗口，实现UIComponent、UIComponentAddtion接口
type WinForm struct {
	// UIAttr 嵌套UI属性
	UIAttr
	// Components 容器的组件列表
	Components []UIComponent
}

// GetUIControlName 获取控件名字
func (window *WinForm) GetUIControlName() string {
	return "WinForm"
}

// GetConcreteUIControlName 获取控件具体名字
func (window *WinForm) GetConcreteUIControlName() string {
	return window.Name
}

// PrintUIComponent 打印UI组件
func (window *WinForm) PrintUIComponent() {
	client.printContainer(window, window)
}

// AddUIComponent 添加UI组件
func (window *WinForm) AddUIComponent(component UIComponent) {
	window.Components = append(window.Components, component)
}

// AddUIComponents 添加UI组件列表
func (window *WinForm) AddUIComponents(components []UIComponent) {
	window.Components = append(window.Components, components...)
}

// GetUIComponentList 获取UI组件列表
func (window *WinForm) GetUIComponentList() []UIComponent {
	return window.Components
}

// Frame 框架，实现UIComponent、接口
type Frame struct {
	// UIAttr 嵌套UI属性
	UIAttr
	// Components 容器的组件列表
	Components []UIComponent
}

// GetUIControlName 获取控件名字
func (frame *Frame) GetUIControlName() string {
	return "Frame"
}

// GetConcreteUIControlName 获取控件具体名字
func (frame *Frame) GetConcreteUIControlName() string {
	return frame.Name
}

// PrintUIComponent 打印UI组件
func (frame *Frame) PrintUIComponent() {
	client.printContainer(frame, frame)
}

// AddUIComponent 添加UI组件
func (frame *Frame) AddUIComponent(component UIComponent) {
	frame.Components = append(frame.Components, component)
}

// AddUIComponents 添加UI组件列表
func (frame *Frame) AddUIComponents(components []UIComponent) {
	frame.Components = append(frame.Components, components...)
}

// GetUIComponentList 获取UI组件列表
func (frame *Frame) GetUIComponentList() []UIComponent {
	return frame.Components
}
