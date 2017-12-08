package decSecondweek

// date:2017-12-05

// 题目：Divide Two Integers
// 不用乘法，除法和mod运算符来计算两个整数的除法。 如果溢出，则返回MAX_INT。
// Divide two integers without using multiplication, division and mod operator.
// If it is overflow, return MAX_INT.

// Time Limit Exceeded 思路:用被除数累加法来判断,但是如果除数太大会超时
func divide(dividend, divisor int) int {
	mutiple := 0
	signCount := 0
	if dividend < 0 {
		signCount++
		dividend = -dividend
	}
	if divisor < 0 {
		signCount++
		divisor = -divisor
	}

	maxInt32 := 1<<31 - 1
	sum := divisor

	for sum <= dividend {
		mutiple++
		sum += divisor
	}
	if mutiple > maxInt32 {
		mutiple = maxInt32
	}
	if signCount == 1 {
		mutiple = -mutiple
	}
	return mutiple
}

// 使用递归,方法很巧妙,但是不太懂.
func divideByRecursion(dividend, divisor int) int {
	maxUint32 := 1<<31 - 1
	mutiple := divideRecur(dividend, divisor)

	if mutiple > maxUint32 {
		mutiple = maxUint32
	}

	return mutiple
}

func divideRecur(dividend, divisor int) int {
	// the sign
	isNegative := dividend < 0 != (divisor < 0)

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	if dividend < divisor {
		return 0
	}

	mutiple := 1
	sum := divisor
	// TODO: 巧妙之处,1,2,4,8...最后会dividend<divisor退出
	for sum+sum <= dividend {
		mutiple += mutiple
		sum += sum
	}
	result := mutiple + divideRecur((dividend-sum), divisor)
	if isNegative {
		result = -result
	}
	return result
}
