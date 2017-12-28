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
