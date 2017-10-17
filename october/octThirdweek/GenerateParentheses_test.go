package octThirdweek

import (
	"reflect"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	if !reflect.DeepEqual(GenerateParentheses(3), 0) {
		t.Errorf("测试失败！ %v", GenerateParentheses(3))
	}
}
