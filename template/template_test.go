package template

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("template: my", MyAskForLeaveTemplate)
	t.Run("template: tom", TomAskForLeaveTemplate)
}

func MyAskForLeaveTemplate(t *testing.T) {
	request := &MyAskForLeaveRequest{}
	company := Company{request}
	company.AskLeave()
}

func TomAskForLeaveTemplate(t *testing.T) {
	request := &TomAskForLeaveRequest{}
	company := Company{request}
	company.AskLeave()
}
