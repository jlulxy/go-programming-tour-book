package errcode

var (
	ErrorGetTagListFail = NewError(2001001, "获取标签列表失败")
	ErrorCreateTagFail  = NewError(2001002, "创建标签失败")
	ErrorUpdateTagFail  = NewError(2001003, "更新标签失败")
	ErrorDeleteTagFail  = NewError(2001004, "删除标签失败")
	ErrorCountTagFail   = NewError(2001005, "统计标签失败")
)
