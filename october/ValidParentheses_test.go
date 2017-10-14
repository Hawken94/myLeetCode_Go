package october

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	s := "(){}[]"
	if !reflect.DeepEqual(IsValid(s), 0) {
		t.Errorf("测试失败！ %v", IsValid(s))
	}
}
