package xerr

import (
	"fmt"
)

type CodeErr struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeErr) GetErrCode() uint32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e CodeErr) GetErrMsg() string {
	return e.errMsg
}
func (e CodeErr) Error() string {
	return fmt.Sprintf("ErrCode:%d, ErrMsg:%s", e.errCode, e.errMsg)
}
func NewErrCode(errCode uint32) *CodeErr {
	return &CodeErr{errCode: errCode, errMsg: MapErrMsg(errCode)}
}
func NewErrMsg(errMsg string) *CodeErr {
	return &CodeErr{errCode: SEVER_COMMON_ERROR, errMsg: errMsg}
}
