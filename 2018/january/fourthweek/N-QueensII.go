package fourthweek

// date:2018-1-27
// 题目: N 皇后2 (N-Queens II)
// tag:回溯 递归

// 与 N 皇后类似,比它简单,只需要输出符合的结果,
import (
	"fmt"
)

var count int

// TODO: 本地调试没有错误,但是 LeetCode 说n=2的时候会出错
func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}
	col := make([]int, n) // i 为行,值为列的数组
	fmt.Println(col)
	dfs2(col, 0, n)
	return count
}
func dfs2(col []int, cur, n int) {
	if cur == n {
		count++
	} else {
		for i := 0; i < n; i++ {
			col[cur] = i
			if validate2(col, cur) {
				dfs2(col, cur+1, n)
			}
		}
	}

}
func validate2(col []int, cur int) bool {
	for i := 0; i < cur; i++ {
		if col[i] == col[cur] || col[i]+i == col[cur]+cur || col[i]+cur == col[cur]+i {
			return false
		}
	}
	return true
}
