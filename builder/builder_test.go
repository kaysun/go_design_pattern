package builder

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("builder-maxTotalNotSet: ", maxTotalNotSet)
	t.Run("builder-minIdleMoreThanMaxIdle: ", minIdleMoreThanMaxIdle)
	t.Run("builder-buildSuccess: ", buildSuccess)
}

func maxTotalNotSet(t *testing.T) {
	// 此为示例代码，在实践中，不应将建造者与使用者放在一个包内，所以这里假设不能直接访问resourcePoolConfig
	build := &ResourcePoolBuild{}
	build.SetName("mysql连接池")
	build.SetMaxIdle(10)
	errCode := build.Build()
	fmt.Println(fmt.Sprintf("建造资源池配置：%s", errMap[errCode]))
}

func minIdleMoreThanMaxIdle(t *testing.T) {
	// 此为示例代码，在实践中，不应将建造者与使用者放在一个包内，所以这里假设不能直接访问resourcePoolConfig
	build := &ResourcePoolBuild{}
	build.SetName("mysql连接池")
	build.SetMaxIdle(10)
	build.SetMinIdle(20)
	build.SetMaxTotal(50)
	errCode := build.Build()
	fmt.Println(fmt.Sprintf("建造资源池配置：%s", errMap[errCode]))
}

func buildSuccess(t *testing.T) {
	// 此为示例代码，在实践中，不应将建造者与使用者放在一个包内，所以这里假设不能直接访问resourcePoolConfig
	build := &ResourcePoolBuild{}
	build.SetName("mysql连接池")
	build.SetMaxIdle(10)
	build.SetMinIdle(5)
	build.SetMaxTotal(50)
	errCode := build.Build()
	fmt.Println(fmt.Sprintf("建造资源池配置：%s", errMap[errCode]))
}
