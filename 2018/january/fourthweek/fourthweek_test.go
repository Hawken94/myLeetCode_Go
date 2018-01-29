package fourthweek

import (
	"testing"
)

func TestNQueens(t *testing.T) {
	t.Errorf("sovleNQueens :%v \n", solveNQueens(0))
}
func TestNQueensII(t *testing.T) {
	t.Errorf("totalNQueens :%v \n", totalNQueens(2))
}

func TestJumpGame(t *testing.T) {
	var tests = []struct {
		nums []int
	}{
		{[]int{2, 3, 1, 1, 4}},
		{[]int{3, 2, 1, 0, 4}},
	}
	for _, test := range tests {
		t.Errorf("canJump result :%v \n", canJump(test.nums))
	}
}

func TestMerge(t *testing.T) {
	var list = IntervalList{
		{1, 3},
		{0, 3},
		{2, 3},
		{4, 6},
	}
	t.Errorf("merge :%v \n", merge(list))
}
