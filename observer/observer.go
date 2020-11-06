// package observer 观察者模式，以两种类型的观察者（可以有多个观察者）为例，向被观察者添加观察者，观察者完成操作后，通知观察者
package observer

import "fmt"

// Observer 观察者
type Observer interface {
	//notify 通知
	notify(sub Subject)
}

// BaseObserver 基础观察者
type BaseObserver struct {
	Name string
}

// TeacherObserver 老师观察者
type TeacherObserver struct {
	BaseObserver
}

// ParentObserver 家长观察者
type ParentObserver struct {
	BaseObserver
}

// notify 老师观察者，实现Observer接口
func (to TeacherObserver) notify(sub Subject) {
	fmt.Println(to.Name + "老师收到了作业")
}

// notify 家长观察者，实现Observer接口
func (to ParentObserver) notify(sub Subject) {
	fmt.Println(to.Name + "(家长)收到了作业")
}
