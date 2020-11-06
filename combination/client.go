package combination

import "fmt"

// PrintClient 打印客户端
type PrintClient struct{}

// printContainer 打印容器控件
func (client PrintClient) printContainer(component UIComponent, componentAddtion UIComponentAddtion) {
	client.printCurrentControl(component)
	for _, v := range componentAddtion.GetUIComponentList() {
		v.PrintUIComponent()
	}
}

// printCurrentControl 打印当前控件
func (client PrintClient) printCurrentControl(component UIComponent) {
	fmt.Println(fmt.Sprintf("print %s(%s)", component.GetUIControlName(), component.GetConcreteUIControlName()))
}
