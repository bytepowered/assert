//go:build no_assert

package assert

func SetEnabled(v bool) {

}

func MustTrue(expectTrue bool, message string, msgArgs ...any) {
}

func MustFalse(expectFalse bool, message string, args ...any) {

}

func MustEmpty(expectEmptyStr string, message string, args ...any) {

}

func MustNotEmpty(expectNotEmptyStr string, message string, args ...any) {

}

func MustNil(expectNil any, message string, args ...any) {
}

func MustNotNil(expectNotNil any, message string, args ...any) {

}
