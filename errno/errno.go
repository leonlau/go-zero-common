package errno

import "strings"

var (
	defaultError error = ErrInternal
)

var (
	OK                      = NewError(200, "OK")
	ErrInternal             = NewError(1000, "服务端错误")
	ErrUnauthorized         = NewError(1001, "未认证")
	ErrForbidden            = NewError(1002, "没有权限")
	ErrNotFound             = NewError(1003, "未找到记录")
	ErrInvalidArgument      = NewError(1004, "参数错误")
	ErrAlreadyExists        = NewError(1005, "已存在")
	ErrAborted              = NewError(1006, "已终止")
	ErrUnavailable          = NewError(1007, "服务不可用")
	ErrAuthenticationFailed = NewError(1008, "认证失败")
)

func NewError(code int, msg string, descs ...string) error {
	desc := ""
	if len(descs) != 0 {
		desc = strings.Join(descs, ",")
	}

	return &BaseResp{
		Code: code,
		Msg:  msg,
		Desc: desc,
	}
}
