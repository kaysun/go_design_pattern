package observer

import "fmt"

//Subject 被观察者接口
type Subject interface {
	//通知观察者
	NotifyObservers()

	//添加观察者
	AddObserver(observer Observer)

	//批量添加观察者
	AddObservers(observers ...Observer)

	//移除观察者
	RemoveObserver(observer Observer)

	//移除全部观察者
	RemoveAllObservers()
}

// StudentSubject 学生被观察者，实现Subject接口
type StudentSubject struct {
	// observers 观察者们
	observers []Observer
}

// NotifyObservers 通知观察者 注：需要使用指针接收者，不然s.observers的值会为空
func (ss *StudentSubject) NotifyObservers() {
	fmt.Println("学生写完了作业，通知观察者们")
	for _, o := range ss.observers {
		//通知观察者
		o.notify(ss)
	}
}

// AddObserver 添加观察者
func (ss *StudentSubject) AddObserver(observer Observer) {
	ss.observers = append(ss.observers, observer)
}

// AddObservers 添加观察者列表
func (ss *StudentSubject) AddObservers(observers ...Observer) {
	ss.observers = append(ss.observers, observers...)
}

// RemoveObserver 移除观察者
func (ss *StudentSubject) RemoveObserver(observer Observer) {
	for i := 0; i < len(ss.observers); i++ {
		if ss.observers[i] == observer {
			ss.observers = append(ss.observers[:i], ss.observers[i+1:]...)
		}
	}
}

// RemoveAllObservers 移除全部观察者
func (ss *StudentSubject) RemoveAllObservers() {
	ss.observers = ss.observers[:0]
}
