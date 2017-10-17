package sepEighthweek

import (
	"reflect"
	"testing"
)

func TestEightweek(t *testing.T) {
	// Integer to Roman
	if !reflect.DeepEqual(IntToRoman(1996), "") {
		t.Errorf("测试失败！ %v", IntToRoman(1996))
	}

	// Roman to Integer
	if !reflect.DeepEqual(RomanToInt("MCMXCVI"), 1996) {
		t.Errorf("测试失败！ %v", RomanToInt("MCMXCVI"))
	}

	// Longest Common Prefix
	if !reflect.DeepEqual(LongestCommonPrefix([]string{"aa", "aa"}), "") {
		t.Errorf("测试失败！ %v", LongestCommonPrefix([]string{"aa", "aa"}))
	}

	// 3Sum Closest
	if !reflect.DeepEqual(ThreeSumClosest([]int{1, 1, 1, 0}, 100), 0) {
		t.Errorf("测试失败！ %v", ThreeSumClosest([]int{1, 1, 1, 0}, 100))
	}
}
