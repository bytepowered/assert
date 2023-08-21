package assert

import (
	"fmt"
	"os"
	"reflect"
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

func MustTrue(isTrue bool, message string, msgArgs ...any) {
	if enabled && !isTrue {
		panic(fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, msgArgs...)))
	}
}

func MustFalse(isFalse bool, message string, args ...any) {
	MustTrue(!isFalse, message, args...)
}

func MustEmpty(str string, message string, args ...any) {
	MustTrue("" == str, message, args...)
}

func MustNotEmpty(str string, message string, args ...any) {
	MustTrue("" != str, message, args...)
}

// MustNil 对输入值进行断言，期望为Nil(包含nil和值nil情况)；
// 如果输入值为非Nil，断言将触发panic，抛出错误消息（消息模板）。
func MustNil(v any, message string, args ...any) {
	if enabled && !IsNil(v) {
		var perr error
		if err, ok := v.(error); ok {
			perr = fmt.Errorf(ErrMsgPrefix+"%s %w", fmt.Sprintf(message, args...), err)
		} else {
			perr = fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, args...))
		}
		panic(perr)
	}
}

// MustNotNil 对输入值进行断言，期望为非Nil；
// 如果输入值为Nil值（包括nil和值为Nil情况），断言将触发panic，抛出错误消息（消息模板）。
func MustNotNil(v any, message string, args ...any) {
	if enabled && IsNil(v) {
		panic(fmt.Errorf(ErrMsgPrefix+"%s", fmt.Sprintf(message, args...)))
	}
}

// IsNil 判断输入值是否为Nil值（包括：nil、类型非Nil但值为Nil），用于检查类型值是否为Nil。
// 只针对引用类型判断有效，任何数值类型、结构体非指针类型等均为非Nil值。
func IsNil(v any) bool {
	if nil == v {
		return true
	}
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map,
		reflect.Interface, reflect.Slice,
		reflect.Ptr, reflect.UnsafePointer:
		return value.IsNil()
	}
	return false
}
