package secondweek

import (
	"testing"
)

func TestSearchForARange(t *testing.T) {
	arr := []int{5, 7, 8, 8, 8, 10}
	result := searchRange(arr, 10)
	t.Errorf("结果为：%v \n", result)
}
