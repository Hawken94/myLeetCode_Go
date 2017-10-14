package october

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	num := []int{-1, 0, 1, 2, -1, -4}
	if !reflect.DeepEqual(FourSum(num, -1), [][]int{{-3, 0, 1, 2}}) {
		t.Errorf("测试失败！ %v", FourSum(num, -1))
	}
}
