//go:build !no_assert

package assert

import (
	"fmt"
	"os"
)

const (
	ErrMsgPrefix = "CRITICAL:ASSERT:"
)

var (
	// 允许通过环境变量禁用断言
	enabled = func() bool {
		v, ok := os.LookupEnv("ASSERT_DISABLED")
		return !ok || "true" != v
	}()
)

func SetEnabled(v bool) {
	enabled = v
}

func MustTrue(expectTrue bool, message string, msgArgs ...any) {
	if enabled && !expectTrue {
		panic(fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, msgArgs...)))
	}
}

func MustFalse(expectFalse bool, message string, args ...any) {
	MustTrue(!expectFalse, message, args...)
}

func MustEmpty(expectEmptyStr string, message string, args ...any) {
	MustTrue("" == expectEmptyStr, message, args...)
}

func MustNotEmpty(expectNotEmptyStr string, message string, args ...any) {
	MustTrue("" != expectNotEmptyStr, message, args...)
}

// MustNil 对输入值进行断言，期望为Nil(包含nil和值nil情况)；
// 如果输入值为非Nil，断言将触发panic，抛出错误消息（消息模板）。
func MustNil(expectNil any, message string, args ...any) {
	if enabled && !IsNil(expectNil) {
		var perr error
		if err, ok := expectNil.(error); ok {
			perr = fmt.Errorf(ErrMsgPrefix+"%s %w", fmt.Sprintf(message, args...), err)
		} else {
			perr = fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, args...))
		}
		panic(perr)
	}
}

// MustNotNil 对输入值进行断言，期望为非Nil；
// 如果输入值为Nil值（包括nil和值为Nil情况），断言将触发panic，抛出错误消息（消息模板）。
func MustNotNil(expectNotNil any, message string, args ...any) {
	if enabled && IsNil(expectNotNil) {
		panic(fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, args...)))
	}
}
