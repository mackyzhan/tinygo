package session

import (
	"errors"
	"fmt"
)

// 错误信息
type Error string

// Format 格式化错误信息并生成新的错误信息
func (this Error) Format(data ...interface{}) Error {
	return Error(fmt.Sprintf(string(this), data...))
}

// Error 生成error类型
func (this Error) Error() error {
	return errors.New(string(this))
}

// String 返回错误字符串描述
func (this Error) String() string {
	return string(this)
}

// 错误码
const (
	ErrorInvalidSessionProvider Error = "ErrorInvalidSessionProvider(S10000):无效的Session提供器"
	ErrorInvalidSessionKind     Error = "ErrorInvalidSessionKind(S10010):无效的Session类型(%s)"
)
