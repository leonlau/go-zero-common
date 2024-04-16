package errno

import (
	"context"
	"fmt"
	"reflect"
)

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Desc string `json:"-"`
	Data any    `json:"data,omitempty"`
}

// type ErrorResp struct {
// 	Code int    `json:"code"`
// 	Msg  string `json:"message"`
// 	Desc string `json:"-"`
// }

func (e *BaseResp) Error() string {
	return fmt.Sprintf("%s:%s", e.Msg, e.Desc)
}

var ErrorHandler = func(ctx context.Context, err error) (int, any) {

	e, ok := err.(*BaseResp)
	if ok {
		// return 200, &ErrorResp{
		// 	Code: e.Code,
		// 	Msg:  e.Msg,
		// 	// Desc: e.Desc,
		// }
		return 200, e
	}
	return 200, defaultError
}

var OkHandler = func(ctx context.Context, data any) any {

	if isEmpty(data) {
		data = nil
	}

	return &BaseResp{
		Code: OK.(*BaseResp).Code,
		Msg:  OK.(*BaseResp).Msg,
		Desc: OK.(*BaseResp).Desc,
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
