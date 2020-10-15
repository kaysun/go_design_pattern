package filter

import "testing"

func Test(t *testing.T) {
	t.Run("comment_filter1: ", CommentFilter1)
	t.Run("comment_filter2: ", CommentFilter2)
	t.Run("comment_filter3: ", CommentFilter3)
	t.Run("up_filter1: ", UpFilter1)
	t.Run("up_filter2: ", UpFilter2)
}

func CommentFilter1(t *testing.T) {
	//构造过滤器信息
	info := CommentInfo{
		Content:   "",
		ArticleID: 1,
	}
	//执行过滤器
	var commentFilter CommentParamFilter
	commentFilter = CommentParamFilter{info}
	commentFilter.DoParamFilter()
}

func CommentFilter2(t *testing.T) {
	//构造过滤器信息
	info := CommentInfo{
		Content:   "太可怕了",
		ArticleID: 0,
	}
	//执行过滤器
	var commentFilter CommentParamFilter
	commentFilter = CommentParamFilter{info}
	commentFilter.DoParamFilter()
}

func CommentFilter3(t *testing.T) {
	//构造过滤器信息
	info := CommentInfo{
		Content:   "    ",
		ArticleID: 1234,
	}
	//执行过滤器
	var commentFilter CommentParamFilter
	commentFilter = CommentParamFilter{info}
	commentFilter.DoParamFilter()
}

func UpFilter1(t *testing.T) {
	//构造过滤器信息
	info := UpInfo{
		CommenID: 0,
		UserID:   6379,
	}
	//执行过滤器
	var upFilter UpParamFilter
	upFilter = UpParamFilter{info}
	upFilter.DoParamFilter()
}

func UpFilter2(t *testing.T) {
	//构造过滤器信息
	info := UpInfo{
		CommenID: 5678,
		UserID:   0,
	}
	//执行过滤器
	var upFilter UpParamFilter
	upFilter = UpParamFilter{info}
	upFilter.DoParamFilter()
}