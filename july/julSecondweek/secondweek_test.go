package julSecondweek

import (
	"reflect"
	"testing"
)

func TestSecondweek(t *testing.T) {
	// Longest Substring Without Repeating Characters
	str := "pwwkew"
	if !reflect.DeepEqual(LengthOfLongestSubstrings(str), 3) {
		t.Fatalf("测试失败！ %v", LengthOfLongestSubstrings(str))
	}

	// Median of Two Sorted Arrays
	var num1 = []int{1, 4}
	var num2 = []int{2, 3, 4}
	if !reflect.DeepEqual(FindMedianSortedArrays(num1, num2), float64(3)) {
		t.Fatalf("测试失败！ %v", FindMedianSortedArrays(num1, num2))
	}
}
