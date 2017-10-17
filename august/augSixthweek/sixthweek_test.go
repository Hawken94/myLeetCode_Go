package augSixthweek

import (
	"reflect"
	"testing"
)

func TestSixweek(t *testing.T) {
	// ZigZag Conversion
	if !reflect.DeepEqual(Convert("PAYPALISHIRING", 3), "") {
		t.Errorf("测试失败！ %v", Convert("PAYPALISHIRING", 3))
	}

	// Container With Most Water
	if !reflect.DeepEqual(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 1) {
		t.Errorf("测试失败！ %v", MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	}
}
