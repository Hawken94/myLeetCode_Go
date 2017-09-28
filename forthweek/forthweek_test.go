package forthweek

import (
	"reflect"
	"testing"
)

func TestForthweek(t *testing.T) {
	// Reverse Integer
	if !reflect.DeepEqual(Reverse(2147483647), 74638474121) {
		t.Errorf("测试失败！ %v", Reverse(2147483647))
	}

	// Implement atoi to convert a string to an integer
	if !reflect.DeepEqual(MyAtoi("9223372036854775809"), 9223372) {
		t.Errorf("测试失败！ %v", MyAtoi("9223372036854775809"))
	}

	// Palindrome Number
	if !IsPalindrome(123) {
		t.Errorf("测试失败！ %v", IsPalindrome(123))
	}
}
