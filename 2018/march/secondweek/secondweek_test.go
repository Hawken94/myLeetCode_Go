package secondweek

import (
	"testing"
)

func TestRotateList(t *testing.T) {
	var tests = []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{1, 2}, 1},
	}
	for _, test := range tests {
		t.Errorf("rotate list now :%v \n", rotate2(test.nums, test.k))
	}
}
