package errcode

var (
	Success                = NewError(0, "成功")
	ServerError            = NewError(100000, "服务内部错误")
	InvalidParams          = NewError(100001, "入参错误")
	NotFound               = NewError(100002, "找不到")
	UnauthorizedAuthExit   = NewError(100003, "鉴权失败，鉴权参数缺失")
	UnauthorizedTokenError = NewError(100004, "鉴权失败，token错误")
	TimeOut                = NewError(100005, "超时")
	TooManyRequest         = NewError(100006, "连接过多")
)
