package combination

// Picture 图片，实现UIComponent接口
type Picture struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (picture Picture) GetUIControlName() string {
	return "Picture"
}

// GetConcreteUIControlName 获取控件具体名字
func (picture Picture) GetConcreteUIControlName() string {
	return picture.Name
}

// PrintUIComponent 打印UI组件
func (picture Picture) PrintUIComponent() {
	client.printCurrentControl(picture)
}

// Button 按钮，实现UIComponent接口
type Button struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (button Button) GetUIControlName() string {
	return "Button"
}

// GetConcreteUIControlName 获取控件具体名字
func (button Button) GetConcreteUIControlName() string {
	return button.Name
}

// PrintUIComponent 打印UI组件
func (button Button) PrintUIComponent() {
	client.printCurrentControl(button)
}

// Label 标签，实现UIComponent接口
type Label struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (label Label) GetUIControlName() string {
	return "Label"
}

// GetConcreteUIControlName 获取控件具体名字
func (label Label) GetConcreteUIControlName() string {
	return label.Name
}

// PrintUIComponent 打印UI组件
func (label Label) PrintUIComponent() {
	client.printCurrentControl(label)
}

// TextBox 文本框，实现UIComponent接口
type TextBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (textBox TextBox) GetUIControlName() string {
	return "TextBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (textBox TextBox) GetConcreteUIControlName() string {
	return textBox.Name
}

// PrintUIComponent 打印UI组件
func (textBox TextBox) PrintUIComponent() {
	client.printCurrentControl(textBox)
}

// PassWordBox 密码框，实现UIComponent接口
type PassWordBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (passWordBox PassWordBox) GetUIControlName() string {
	return "PassWordBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (passWordBox PassWordBox) GetConcreteUIControlName() string {
	return passWordBox.Name
}

// PrintUIComponent 打印UI组件
func (passWordBox PassWordBox) PrintUIComponent() {
	client.printCurrentControl(passWordBox)
}

// CheckBox 复选框，实现UIComponent接口
type CheckBox struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (checkBox CheckBox) GetUIControlName() string {
	return "CheckBox"
}

// GetConcreteUIControlName 获取控件具体名字
func (checkBox CheckBox) GetConcreteUIControlName() string {
	return checkBox.Name
}

// PrintUIComponent 打印UI组件
func (checkBox CheckBox) PrintUIComponent() {
	client.printCurrentControl(checkBox)
}

// LinkLabel 关联的标签，实现UIComponent接口
type LinkLabel struct {
	// UIAttr 嵌套UI属性
	UIAttr
}

// GetUIControlName 获取控件名字
func (linkLabel LinkLabel) GetUIControlName() string {
	return "LinkLabel"
}

// GetConcreteUIControlName 获取控件具体名字
func (linkLabel LinkLabel) GetConcreteUIControlName() string {
	return linkLabel.Name
}

// PrintUIComponent 打印UI组件
func (linkLabel LinkLabel) PrintUIComponent() {
	client.printCurrentControl(linkLabel)
}
