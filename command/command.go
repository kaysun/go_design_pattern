// package command 命令模式，引入调用者和接收者实现解耦
package command

import "fmt"

// Command 命令接口
type Command interface {
	// Execute 执行命令
	Execute()
}

// CreateCommand 创建命令，实现Command接口
type CreateCommand struct {
	// receiver 接收者接口对象
	receiver Receiver
}

// UpdateCommand 更新命令，实现Command接口
type UpdateCommand struct {
	// receiver 接收者接口对象
	receiver Receiver
}

// Execute 创建命令实现执行命令方法
func (command CreateCommand)Execute() {
	command.receiver.Action()
}

// Execute 修改命令实现执行命令方法
func (command UpdateCommand)Execute() {
	command.receiver.Action()
}

// Receiver 接收者接口
type Receiver interface {
	// Action 动作
	Action()
}

// CreateReceiver 创建接收者，实现Receiver接口
type CreateReceiver struct {}

// UpdateReceiver 更新接收者，实现Receiver接口
type UpdateReceiver struct {}

func (receiver CreateReceiver) Action() {
	fmt.Println("执行了创建")
}

func (receiver UpdateReceiver) Action() {
	fmt.Println("执行了修改")
}

// Invoker 调用者
type Invoker struct {
	// command 命令接口对象
	command Command
}

// Call 调用者调用执行命令方法
func (invoker Invoker) Call() {
	invoker.command.Execute()
}