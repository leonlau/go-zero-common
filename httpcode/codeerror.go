package httpcode

import (
	"context"
	"fmt"
)

type BaseResp struct {
	CodeError `json:",inline"`
	Data      any `json:"data,omitempty"`
}
type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Desc string `json:"desc,omitempty"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("%s:%s", e.Msg, e.Desc)
}

func NewOk(data any) *BaseResp {
	if data == nil {
		data = struct{}{}
	}
	return &BaseResp{
		CodeError: CodeError{
			Code: defaultOKCode,
			Msg:  defaultOKMsg,
		},
		Data: data,
	}
}

var ErrorHandler = func(ctx context.Context, err error) (int, any) {
	e, ok := err.(*CodeError)
	if ok {
		return 200, e
	}
	return 200, CodeError{
		Code: defaultErrorCode,
		Msg:  getErrorMsg(defaultErrorCode),
		// Desc: err.Error(),
	}
}
