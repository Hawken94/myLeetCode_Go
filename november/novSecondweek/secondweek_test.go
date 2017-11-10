package novSecondweek

import (
	"testing"
)

// func TestMergeKLists(t *testing.T) {

// }

func TestSearchForARange(t *testing.T) {
	arr := []int{5, 7, 8, 8, 8, 10}
	result := searchRange(arr, 10)
	t.Errorf("结果为：%v \n", result)
}
func TestSearchInsert(t *testing.T) {
	t.Errorf("xhk: %v \n", searchInsert([]int{1, 4, 4, 4, 4, 4, 4, 4, 4}, 4))
}
