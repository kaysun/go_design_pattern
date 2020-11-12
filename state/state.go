package state

import "fmt"

// State 用户状态
type State uint8

const (
	// StateNormal 普通用户
	StateNormal State = 1
	// StateVip 普通会员
	StateVip State = 2
)

// ISwitchState 转换状态接口
type ISwitchState interface {
	// PurchaseVip 购买会员
	PurchaseVip()
	// Expire 会员过期
	Expire()
}

// IUser 用户接口
type IUser interface {
	// WatchVideo 看视频
	WatchVideo()
}

// NormalUser 普通用户，实现IUser接口
type NormalUser struct{}

// WatchVideo 看视频
func (user NormalUser) WatchVideo() {
	fmt.Println("看广告中...")
}

// VipUser 会员用户，实现IUser接口
type VipUser struct{}

// WatchVideo 看视频
func (user VipUser) WatchVideo() {
	fmt.Println("您是尊敬的vip用户，已为您跳过120s广告")
}

// User 用户，实现ISwitchState、IUser接口
type User struct {
	// UserState 用户
	UserState IUser
}

// SetUser 设置用户状态
func (user *User) SetUser(userState IUser) {
	user.UserState = userState
}

// Expire 会员过期
func (user *User) Expire() {
	user.UserState = NormalUser{}
}

// PurchaseVip 购买vip会员
func (user *User) PurchaseVip() {
	user.UserState = VipUser{}
}

// WatchVideo 看视频
func (user *User) WatchVideo() {
	if user.UserState == nil {
		// 默认为普通用户
		user.SetUser(NormalUser{})
	}
	user.UserState.WatchVideo()
}
