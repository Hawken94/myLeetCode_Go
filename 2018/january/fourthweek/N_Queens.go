package fourthweek

// date:2018-1-27

// 题目:N-Queens
// N 皇后,不能在同一行,同一列,同一对角线;用'Q' 代替皇后,输出所有结果
// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.

// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
// For example,
// There exist two distinct solutions to the 4-queens puzzle:

func solveNQueens(n int) [][]string {
	board := [][]byte{}
	// 这里要注意一下二维数组赋值的问题,最里面一层要先赋值
	for i := 0; i < n; i++ {
		var temp []byte
		for j := 0; j < n; j++ {
			temp = append(temp, '.')
		}
		board = append(board, temp)
	}
	res := make([][]string, 0)

	dfs(board, 0, &res)
	return res
}

func dfs(board [][]byte, colIndex int, res *[][]string) {

	if colIndex == len(board) {
		// slice是引用类型,这里用临时变量防止结果被修改
		// tempRes := make([]string, 0)
		*res = append(*res, construct(board))

		return
	}
	for i := 0; i < len(board); i++ {
		if validate(board, i, colIndex) {
			board[i][colIndex] = 'Q'
			dfs(board, colIndex+1, res)
			board[i][colIndex] = '.'
		}
	}
}

func validate(board [][]byte, x, y int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < y; j++ {
			if board[i][j] == byte('Q') && (x == i || x+y == i+j || x+j == y+i) {
				return false
			}
		}
	}
	return true
}

func construct(board [][]byte) []string {
	res := make([]string, 0)
	for i := 0; i < len(board); i++ {
		res = append(res, string(board[i]))
	}
	return res
}
