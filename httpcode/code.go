package httpcode

import (
	"strings"
	"sync"
)

var (
	defaultErrorCode int    = 1000
	defaultOKCode    int    = 200
	defaultOKMsg     string = "OK"
)

var mutex *sync.RWMutex
var errorSet map[int]string

func init() {
	mutex = &sync.RWMutex{}
	errorSet = map[int]string{
		1000: "Internal Error",
		1001: "Unauthorized",
		1002: "Forbidden",
		1003: "Not Found",
		1004: "Invalid Argument",
		1005: "Already Exists",
		1006: "Aborted",
		1007: "Unavailable",
		1008: "AuthenticationFailed",
	}
}

func SetDefaultErrorCode(code int) {
	defaultErrorCode = code
}

func SetOKResq(code int, msg string) {
	defaultOKCode = code
	defaultOKMsg = msg
}

func getErrorMsg(code int) string {
	mutex.RLock()
	defer mutex.RUnlock()
	if msg, ok := errorSet[code]; ok {
		return msg
	}

	return errorSet[defaultErrorCode]
}

func AppendErrorMsg(code int, msg string) {
	mutex.Lock()
	defer mutex.Unlock()
	errorSet[code] = msg
}

func GetCodeError(code int, descs ...string) error {
	desc := ""
	if len(descs) != 0 {
		desc = strings.Join(descs, ",")
	}

	return &BaseResp{
		Code: code,
		Msg:  getErrorMsg(code),
		Desc: desc,
	}
}

func NewInternalError(desc ...string) error {
	return GetCodeError(1000, desc...)
}

func NewUnauthorizedError(desc ...string) error {
	return GetCodeError(1001, desc...)
}

func NewForbiddenError(desc ...string) error {
	return GetCodeError(1002, desc...)
}

func NewNotFoundError(desc ...string) error {
	return GetCodeError(1003, desc...)
}
func NewInvalidArgumentError(desc ...string) error {
	return GetCodeError(1004, desc...)
}
func NewAlreadyExistsError(desc ...string) error {
	return GetCodeError(1005, desc...)
}

func NewAbortedError(desc ...string) error {
	return GetCodeError(1006, desc...)
}

func NewUnavailableError(desc ...string) error {
	return GetCodeError(1007, desc...)
}

func NewAuthenticationFailedError(desc ...string) error {
	return GetCodeError(1008, desc...)
}
