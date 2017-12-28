package decFifthweek

// date:2017-12-25

// 题目:Sqrt(x)
// 实现简单的开方算法
/*
Implement int sqrt(int x).
Compute and return the square root of x.
x is guaranteed to be a non-negative integer.
*/
func mySqrt(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	for i := 1; i < n; i++ {
		if i*i <= n && (i+1)*(i+1) > n {
			return i
		}
	}
	return 0
}
