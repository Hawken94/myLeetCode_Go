package decFifthweek

import (
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	// matrix := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	tests := []struct {
		matrix [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			[][]int{{6, 9, 7}},
		},
	}
	for _, test := range tests {
		t.Errorf("spiralOrder :%v \n", spiralOrder(test.matrix))
	}

}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		s string
	}{
		{"hello world  "},
		{"   "},
		{"hello ,world !"},
	}

	for _, test := range tests {
		t.Errorf("lengthOfLastWord :%v\n", lengthOfLastWord(test.s))
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int // 单元测试要规定nums1的cap
		m     int
		nums2 []int
		n     int
	}{
		{
			[]int{1, 8}, 2, []int{3, 7}, 2,
		},
	}
	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		t.Errorf("merge :%v \n", test.nums1)
	}
}
