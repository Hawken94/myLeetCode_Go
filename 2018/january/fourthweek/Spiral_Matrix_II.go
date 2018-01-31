package fourthweek

// date:2018-1-31
// 题目: 螺旋化矩阵2 (Spiral Matrix II)
// lable:螺旋化矩阵

// 给一个整数 n,生成一个1到 n 的平方的矩阵数组,是Spriral Matrix 的逆向
// Given an integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.
// For example,
// Given n = 3,
// You should return the following matrix:
// [
//  [ 1, 2, 3 ],
//  [ 8, 9, 4 ],
//  [ 7, 6, 5 ]
// ]

// 思路和 Spiral Matrix 一模一样
func generateMatrix(n int) [][]int {
	// TODO: 二维切片的赋值要注意数组越界
	var matrix [][]int
	// var matrix [][]int
	if n == 0 {
		return matrix
	}
	// 二维数组使用前要先赋值,要不然会出现数组越界的情况,真傻逼
	for i := 0; i < n; i++ {
		var temp []int
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		matrix = append(matrix, temp)
	}
	rowBegin, colBegin := 0, 0
	rowEnd, colEnd := n-1, n-1
	number := 1
	for rowBegin <= rowEnd && colBegin <= colEnd {
		// Traverse Right
		for i := colBegin; i <= colEnd; i++ {
			matrix[rowBegin][i] = number
			number++
		}
		rowBegin++

		// Traverse Down
		for i := rowBegin; i <= rowEnd; i++ {
			matrix[i][rowEnd] = number
			number++
		}
		colEnd--

		if rowBegin <= rowEnd {
			// Traverse Left
			for i := colEnd; i >= colBegin; i-- {
				matrix[rowEnd][i] = number
				number++
			}
		}
		rowEnd--

		if colBegin <= colEnd {
			// Traverse Up
			for i := rowEnd; i >= rowBegin; i-- {
				matrix[i][colBegin] = number
				number++
			}
		}
		colBegin++

	}

	return matrix
}
