package errorsWrap

import "testing"

const (
	// Wrap.Error() 应该返回的String
	needWrapErrorResult string = "{ ErrorType: UNIT_TEST_TYPE | ErrorMessage: this is a test message }"

	// Wrap.String() 应该返回的String
	needWrapStringResult string = "{ ErrorType: UNIT_TEST_TYPE | ErrorMessage: this is a test message }"
)

// 基本上不会有问题的单元测试
func TestNew(t *testing.T) {
	err := New("This is a test")
	if err == nil {
		t.Fatalf("nil")
		return
	}
}

// 主要封装函数的单元测试
func TestWrap(t *testing.T) {
	err := Wrap(New("UNIT_TEST_TYPE"), "this is a test message")
	if err == nil {
		t.Fatalf("nil")
		return
	}
	if err.Error() != needWrapErrorResult {
		t.Fatalf(err.Error() + " != " + needWrapErrorResult)
		return
	}
	if err.String() != needWrapStringResult {
		t.Fatalf(err.String() + " != " + needWrapStringResult)
		return
	}
}
