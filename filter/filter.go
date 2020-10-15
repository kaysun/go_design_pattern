/*
	package filter 过滤器模式，责任链模式的一种。
	每个filter都需要验证，完全验证通过才会放行，去执行业务逻辑
	一旦有filter验证不通过，则立刻退出
	适合处理参数校验，权限校验，记录日志，准备相关资源等场景
 */
package filter

import (
	"context"
	"fmt"
	"strings"
)

// Context 过滤器上下文
type Context struct {
	// context.Context 嵌套基础库
	context.Context
	// key 每个过滤器的key
	key string
}

// HandleFunc 定义别名：过滤器函数
type HandlerFunc func(ctx Context, params ...interface{}) error

// HandlersChain 过滤器责任链
type HandlersChain []HandlerFunc

// Filter 过滤器接口
type Filter interface {
	DoFilter() error
}

// FilterImpl 实际的过滤器，封装了过滤器上下文、过滤器责任链
type FilterImpl struct {
	// Ctx 过滤器上下文
	Ctx Context
	// Chain 过滤器责任链
	Chain HandlersChain
}

// DoFilter 执行过滤器，FilterImpl实现Filter接口
func (myFilter *FilterImpl) DoFilter() error {
	for _, handler := range myFilter.Chain {
		err := handler(myFilter.Ctx)
		if err != nil {
			fmt.Println(fmt.Sprintf("filter fail key=%s, err = %v", myFilter.Ctx.key, err))
			return err
		}
	}
	return nil
}

//NewFilter 创建过滤器
func NewFilter(ctx Context, handlers ...HandlerFunc) Filter {
	return &FilterImpl{
		Ctx:   ctx,
		Chain: handlers,
	}
}

/**********过滤器使用***********/

//ParamFilter 参数过滤器接口
type ParamFilter interface {
	// DoParamFilter 执行参数过滤
	DoParamFilter() error
}

// CommentParamFilter 评论参数过滤器，实现ParamFilter接口
type CommentParamFilter struct {
	CommentInfo
}

// UpParamFilter 点赞参数过滤器，实现ParamFilter接口
type UpParamFilter struct {
	UpInfo
}

// DoParamFilter 评论参数过滤器，实现参数过滤方法
func (cpf CommentParamFilter) DoParamFilter() error {
	myContext := Context{
		Context:  context.TODO(),
		key: "评论参数过滤器",
	}
	cf := NewFilter(
		myContext,
		validateComment1(myContext, cpf.CommentInfo),
		validateComment2(myContext, cpf.CommentInfo),
	)
	return cf.DoFilter()
}

// validateComment1 评论参数校验方法1
func validateComment1(ctx Context, info CommentInfo) HandlerFunc {
	return func(ctx Context, params ...interface{}) error {
		if info.ArticleID == 0 {
			return fmt.Errorf("当前评论无文章id，不允许评论")
		}
		return nil
	}
}

// validateComment2 评论参数校验方法2
func validateComment2(ctx Context, info CommentInfo) HandlerFunc {
	return func(ctx Context, params ...interface{}) error {
		if len(info.Content) == 0 || len(strings.Trim(info.Content, " ")) == 0 {
			return fmt.Errorf("当前评论无内容，不允许评论")
		}
		return nil
	}
}

// DoParamFilter 点赞参数过滤器，实现参数过滤方法
func (upf UpParamFilter) DoParamFilter() error {
	myContext := Context{
		Context:  context.TODO(),
		key: "点赞参数过滤器",
	}
	cf := NewFilter(
		myContext,
		validateUp1(myContext, upf.UpInfo),
		validateUp2(myContext, upf.UpInfo),
	)
	return cf.DoFilter()
}

// validateUp1 点赞参数校验方法1
func validateUp1(ctx Context, info UpInfo) HandlerFunc {
	return func(ctx Context, params ...interface{}) error {
		if info.CommenID == 0 {
			return fmt.Errorf("当前点赞无评论id，不允许点赞")
		}
		return nil
	}
}

// validateUp2 点赞参数校验方法2
func validateUp2(ctx Context, info UpInfo) HandlerFunc {
	return func(ctx Context, params ...interface{}) error {
		if info.UserID == 0 {
			return fmt.Errorf("当前点赞无用户id，不允许点赞")
		}
		return nil
	}
}

// CommentInfo 评论
type CommentInfo struct {
	// Content 评论内容
	Content string
	// ArticleID 文章ID
	ArticleID uint64
}

// UpInfo 点赞
type UpInfo struct {
	// CommenID 评论ID
	CommenID uint64
	// UserID 点赞人的用户ID
	UserID uint64
}