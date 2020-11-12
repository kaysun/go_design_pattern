package memorandum

import "testing"

func Test(t *testing.T) {
	t.Run("memorandum", Memorandum)
}

func Memorandum(t *testing.T) {
	// 发起者初始状态
	originator := Originator{}
	originator.State = "start"
	originator.print()

	// 设置备忘录
	caretaker := Caretaker{}
	caretaker.memento = originator.createMemento()
	originator.State = "Stage Two"
	originator.print()

	// 恢复为备忘录
	originator.restoreMemento(caretaker.memento)
	originator.print()
}
