package observer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("observer: ", NotifyObservers)
}

func NotifyObservers(t *testing.T) {
	var subject Subject
	subject = &StudentSubject{}

	var mathTeacherObserver, englishTeacherObserver, motherObserver, fatherObserver Observer
	mathTeacherObserver = &TeacherObserver{BaseObserver{"数学"}}
	englishTeacherObserver = &TeacherObserver{BaseObserver{"英语"}}
	motherObserver = &ParentObserver{BaseObserver{"妈妈"}}
	fatherObserver = &ParentObserver{BaseObserver{"爸爸"}}

	//只添加数学老师观察者，那么只有数学老师会收到作业完成通知
	fmt.Println("******只添加数学老师观察者，那么只有数学老师会收到作业完成通知******")
	subject.AddObserver(mathTeacherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//又批量添加了英语老师、妈妈、爸爸，那么数学老师、英语老师、妈妈、爸爸，都会收到作业完成通知
	fmt.Println("******又批量添加了英语老师、妈妈、爸爸，那么数学老师、英语老师、妈妈、爸爸，都会收到作业完成通知******")
	subject.AddObservers(englishTeacherObserver, motherObserver, fatherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//移除了妈妈观察者，那么只有数学老师、英语老师、爸爸，会收到作业完成通知
	fmt.Println("******移除了妈妈观察者，那么只有数学老师、英语老师、爸爸，会收到作业完成通知******")
	subject.RemoveObserver(motherObserver)
	subject.NotifyObservers()
	fmt.Println()

	//移除了所有观察者，那么不会有人收到作业完成通知
	fmt.Println("******移除了所有观察者，那么不会有人收到作业完成通知******")
	subject.RemoveAllObservers()
	subject.NotifyObservers()
	fmt.Println()
}
