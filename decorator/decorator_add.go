// package decorator 装饰者模式，在运行期动态地给对象添加额外的指责，比子类更灵活。第二个功能：用于添加功能的装饰模式。
// 若不想修改原来的接口，则可以使用装饰者。go这种语言，若增加接口的方法，则原来实现接口的结构体都需要增加一个新方法，这样就不符合开闭原则了。
package decorator

import "fmt"

// Booker 书的接口
type Booker interface {
	// Reading 读书
	Reading()
}

// Book 书，实现Booker接口
type Book struct {
}

// Reading 读书，实现Booker接口
func (book Book) Reading() {
	fmt.Println("我正在读书")
}

// Underliner 划线接口，装饰者接口
type Underliner interface {
	// Booker 继承Booker接口
	Booker
	// Underline 划线
	Underline()
}

// NotesTaker 记笔记接口，装饰者接口
type NotesTaker interface {
	// Booker 继承Booker接口
	Booker
	// TakeNotes 记笔记
	TakeNotes()
}

// ConcreteUnderline 具体的划线类，实现Underliner接口
type ConcreteUnderline struct {
	// Booker 书的接口对象
	Booker Booker
}

// ReadingBooks ConcreteUnderline提供读书的方法，包装了Booker接口
func (underline ConcreteUnderline) Reading() {
	underline.Booker.Reading()
}

// Underline 划线，实现Underliner接口
func (underline ConcreteUnderline) Underline() {
	fmt.Println("我正在划线")
}

// ConcreteNotesTake 具体的记笔记类，实现NotesTaker接口
type ConcreteNotesTake struct {
	// Booker 书的接口对象
	Booker Booker
}

// ReadingBooks ConcreteNotesTake提供读书的方法，包装了Booker接口
func (notesTake ConcreteNotesTake) Reading() {
	notesTake.Booker.Reading()
}

// Underline 划线，实现NotesTaker接口
func (notesTake ConcreteNotesTake) TakeNotes() {
	fmt.Println("我正在记笔记")
}
