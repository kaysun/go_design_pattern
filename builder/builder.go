// Package builder 建造者模式
package builder

const (
	// Success 成功
	Success = 0
	// ErrCodeNameRequired 名字必填
	ErrCodeNameRequired = -10001
	// ErrCodeMinIdleMoreThanMaxIdle 最小空闲数超过了最大空闲数
	ErrCodeMinIdleMoreThanMaxIdle = -10002
	// ErrCodeMaxTotalMustMoreThanZero 最大连接数必须大于0
	ErrCodeMaxTotalMustMoreThanZero = -10003
	// ErrCodeMaxIdleMustMoreThanZero 最大空闲数必须大于0
	ErrCodeMaxIdleMustMoreThanZero = -10004
	// ErrCodeMinIdleMustMoreThanZero 最小空闲数必须大于0
	ErrCodeMinIdleMustMoreThanZero = -10005
	// ErrCodeMaxTotalNotSet 最大连接数未设置
	ErrCodeMaxTotalNotSet = -10006
	// ErrCodeMaxIdleNotSet 最大空闲连接数未设置
	ErrCodeMaxIdleNotSet = -10007
	// ErrCodeMinIdleNotSet 最小空闲连接数未设置
	ErrCodeMinIdleNotSet = -10008
)

// errMap 错误信息与错误码的配置
var errMap = map[int16]string{
	Success:                         "ok",
	ErrCodeNameRequired:             "名字必填",
	ErrCodeMinIdleMoreThanMaxIdle:   "最小空闲数超过了最大空闲数",
	ErrCodeMaxTotalMustMoreThanZero: "最大连接数必须大于0",
	ErrCodeMaxIdleMustMoreThanZero:  "最大空闲连接数必须大于0",
	ErrCodeMinIdleMustMoreThanZero:  "最小空闲连接数必须大于0",
	ErrCodeMaxTotalNotSet:           "最大连接数未设置",
	ErrCodeMaxIdleNotSet:            "最大空闲连接数未设置",
	ErrCodeMinIdleNotSet:            "最小空闲连接数未设置",
}

// resourcePoolConfig 资源池配置，不可导出
type resourcePoolConfig struct {
	// name 名字，不可导出，必填
	name string
	// maxTotal 最大连接数，不可导出，必填
	maxTotal uint16
	// maxIdle 最大空闲连接数，不可导出，必填
	maxIdle uint16
	// minIdle 最小空闲连接数，不可导出，不必填，但填里必须大于0
	minIdle uint16
}

// Builder 建造者接口
type Builder interface {
	// Build 创建对象
	Build() int16
}

// ResourcePoolBuild 资源池建造者，实现Builder接口
type ResourcePoolBuild struct {
	// resourcePoolConfig 嵌套资源池配置
	resourcePoolConfig
}

// Build 创建对象
func (build *ResourcePoolBuild) Build() int16 {
	// 这一串的判断可以保证builder里要求必须设置的属性，都能够设置
	if len(build.name) == 0 {
		return ErrCodeNameRequired
	}
	if build.maxTotal == 0 {
		return ErrCodeMaxTotalNotSet
	}
	if build.maxIdle == 0 {
		return ErrCodeMaxIdleNotSet
	}

	// 对设置后的对象，做整体的逻辑性验证
	if build.minIdle > build.maxIdle {
		return ErrCodeMinIdleMoreThanMaxIdle
	}
	return Success
}

// SetName 设置名字，并进行校验
func (build *ResourcePoolBuild) SetName(name string) int16 {
	if len(name) == 0 {
		return ErrCodeNameRequired
	}
	build.name = name
	return Success
}

// SetMaxTotal 设置最大连接数，并进行校验
func (build *ResourcePoolBuild) SetMaxTotal(maxTotal uint16) int16 {
	if maxTotal <= 0 {
		return ErrCodeMaxTotalMustMoreThanZero
	}
	build.maxTotal = maxTotal
	return Success
}

// SetMaxIdle 设置最大空闲连接数，并信息校验
func (build *ResourcePoolBuild) SetMaxIdle(maxIdle uint16) int16 {
	if maxIdle <= 0 {
		return ErrCodeMaxIdleMustMoreThanZero
	}
	build.maxIdle = maxIdle
	return Success
}

// SetMinIdle 设置最小空闲连接数，并进行校验
func (build *ResourcePoolBuild) SetMinIdle(minIdle uint16) int16 {
	if minIdle <= 0 {
		return ErrCodeMinIdleMustMoreThanZero
	}
	build.minIdle = minIdle
	return Success
}
