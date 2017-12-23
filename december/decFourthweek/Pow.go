package decFourthweek

// date:2017-12-23

// 题目:Pow(x, n)
// Implement pow(x, n).

// 递归思想,底数平方,指数减半,当指数为0时退出递归
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)

}
