// package singleton 单例模式
package singleton

import (
	"fmt"
	"sync"
)

var (
	// once 保证只执行一次
	once sync.Once

	// instance 单例对象
	instance *Instance
)

// Instance 单例类型结构体
type Instance struct {
	// Title 标题
	Title string
}

// GetInstance 获取单例对象。 注：必须使用指针对象，否则会在返回时，拷贝一份，对象就不是同一个了。
func GetInstance(title string) *Instance {
	once.Do(func() {
		fmt.Println("执行获取单例对象")
		instance = &Instance{Title: title}
	})
	return instance
}
