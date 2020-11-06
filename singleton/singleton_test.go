package singleton

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("singleton: ", GetSingleton)
}

func GetSingleton(t *testing.T) {
	/* 执行结果：
	   执行获取单例对象
	   instance1 = 0xc0000584a0, name=标题1
	   instance2 = 0xc0000584a0, name=标题1
	*/

	instance1 := GetInstance("标题1")
	fmt.Println(fmt.Sprintf("instance1 = %p, name=%s", instance1, instance1.Title))

	instance2 := GetInstance("标题2")
	fmt.Println(fmt.Sprintf("instance2 = %p, name=%s", instance2, instance2.Title))
}
