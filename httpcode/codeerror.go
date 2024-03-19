package httpcode

import (
	"context"
	"fmt"
	"reflect"
)

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Desc string `json:"desc,omitempty"`
	Data any    `json:"data,omitempty"`
}

type ErrorResp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Desc string `json:"desc,omitempty"`
}

func (e *BaseResp) Error() string {
	return fmt.Sprintf("%s:%s", e.Msg, e.Desc)
}

var ErrorHandler = func(ctx context.Context, err error) (int, any) {

	e, ok := err.(*BaseResp)
	if ok {
		return 200, &ErrorResp{
			Code: e.Code,
			Msg:  e.Msg,
			// Desc: e.Desc,
		}
	}
	return 200, &ErrorResp{
		Code: defaultErrorCode,
		Msg:  getErrorMsg(defaultErrorCode),
		// Desc: err.Error(),
	}
}

var OkHandler = func(ctx context.Context, data any) any {

	if isEmpty(data) {
		data = nil //make([]string, 0)
	}

	return &BaseResp{
		Code: defaultOKCode,
		Msg:  defaultOKMsg,
		Data: data,
	}
}

func isEmpty(obj any) bool {
	if obj == nil {
		return true
	}
	kind := reflect.ValueOf(obj).Kind()
	switch kind {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(obj).IsNil()
	default:
		return false
	}
}
