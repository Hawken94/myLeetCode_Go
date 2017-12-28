package decFifthweek

import (
	"fmt"
)

// date:2017-12-26

// 题目:Spiral Matrix
// 以螺旋方式输出矩阵
/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
For example,
Given the following matrix:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
You should return [1,2,3,6,9,8,7,4,5].
*/

// TODO:这道题还要思考一下
func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)
	if len(matrix) == 0 {
		return result
	}
	fmt.Println(len(matrix), "  ", len(matrix[0])-1)
	rowBegin, colBegin := 0, 0
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		// Traverse Right
		for i := colBegin; i <= colEnd; i++ {
			result = append(result, matrix[rowBegin][i])
		}
		rowBegin++

		// Traverse Down
		for i := rowBegin; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--

		if rowBegin <= rowEnd {
			// Traverse Left
			for i := colEnd; i >= colBegin; i-- {
				result = append(result, matrix[rowEnd][i])
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			// Traverse Up
			for i := rowEnd; i >= rowBegin; i-- {
				result = append(result, matrix[i][colBegin])
			}
		}
		colBegin++
	}

	return result
}
