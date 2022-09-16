package combination

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("combination: ", Combination)
}

// Combination 单元测试
// print WinForm(WINDOW窗口)
// print Picture(LOGO图片)
// print Button(登录)
// print Button(注册)
// print Frame(FRAME1)
// print Label(用户名)
// print TextBox(文本框)
// print Label(密码)
// print PassWordBox(密码框)
// print CheckBox(复选框)
// print TextBox(记住用户名)
// print LinkLabel(忘记密码)
func Combination(t *testing.T) {
	window := &WinForm{UIAttr: UIAttr{Name: "WINDOW窗口"}}
	picture := &Picture{UIAttr{Name: "LOGO图片"}}
	loginButton := &Button{UIAttr{Name: "登录"}}
	registerButton := &Button{UIAttr{Name: "注册"}}
	frame := &Frame{UIAttr: UIAttr{Name: "FRAME1"}}
	userLable := &Label{UIAttr{Name: "用户名"}}
	textBox := &TextBox{UIAttr{Name: "文本框"}}
	passwordLable := &Label{UIAttr{Name: "密码"}}
	passwordBox := &PassWordBox{UIAttr{Name: "密码框"}}
	checkBox := &CheckBox{UIAttr{Name: "复选框"}}
	rememberUserTextBox := &TextBox{UIAttr{Name: "记住用户名"}}
	linkLable := &LinkLabel{UIAttr{Name: "忘记密码"}}

	window.AddUIComponents([]UIComponent{picture, loginButton, registerButton, frame})
	frame.AddUIComponents([]UIComponent{userLable, textBox,
		passwordLable, passwordBox, checkBox, rememberUserTextBox, linkLable})

	window.PrintUIComponent()
}
