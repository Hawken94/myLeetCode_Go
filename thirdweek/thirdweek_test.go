package thirdweek

import (
	"reflect"
	"testing"
)

func TestThirdweek(t *testing.T) {
	// LongestPalindrome
	if !reflect.DeepEqual(LongestPalindrome("caad"), "aa") {
		t.Fatalf("测试失败！ %v", LongestPalindrome("caad"))
	}
}
