package julFirstweek

import (
	"reflect"
	"testing"
)

func TestFirstweek(t *testing.T) {
	// 2Sum
	var nums = []int{2, 7, 3, 6, 11, 15}
	target := 9
	if !reflect.DeepEqual(TwoSumByMap(nums, target), []int{0, 1}) {
		t.Fatalf("测试失败！ %v", TwoSumByMap(nums, target))
	}

	// 3Sum
	var num = IntSlice{-1, -1, -1, 0, 2, 1, -2, 2, 4, 4}
	if !reflect.DeepEqual(ThreeSum(num), [][]int{{-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}) {
		t.Fatalf("测试失败！ %v", ThreeSum(num))
	}

	// Add Two Numbers
	// 失败了，以后再说
}
