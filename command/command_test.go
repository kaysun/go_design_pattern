package command

import "testing"

func Test(t *testing.T) {
	t.Run("create_command: ", CreateCommandExecute)
	t.Run("update_command: ", UpdateCommandExecute)
}

func CreateCommandExecute(t *testing.T) {
	var receiver Receiver
	receiver = &CommentCreateReceiver{}

	var command Command
	command = &CreateCommand{receiver: receiver}

	invoker := Invoker{command: command}
	invoker.Call()
}

func UpdateCommandExecute(t *testing.T) {
	var receiver Receiver
	receiver = &CommentUpdateReceiver{}

	var command Command
	command = &UpdateCommand{receiver: receiver}

	invoker := Invoker{command: command}
	invoker.Call()
}
