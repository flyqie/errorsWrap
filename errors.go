// 一个简单的go errors包装库
package errorsWrap

import (
	goErrors "errors"
	"fmt"
)

type Errors struct {
	// 错误信息
	Messages string
	// 错误类别
	ErrType error
}

// 错误输出
func (e *Errors) Error() string {
	return e.String()
}

// 错误输出
func (e *Errors) String() string {
	return fmt.Sprintf("{ ErrorType: %s | ErrorMessage: %s }", e.ErrType.Error(), e.Messages)
}

// 错误封装
func Wrap(err error, message string) *Errors {
	return &Errors{
		Messages: message,
		ErrType:  err,
	}
}

// 继承自go自带的errors.New函数
func New(text string) error {
	return goErrors.New(text)
}
