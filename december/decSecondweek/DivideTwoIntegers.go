package decSecondweek

// Time Limit Exceeded
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

	maxUint32 := 1<<31 - 1
	sum := divisor

	for sum <= dividend {
		mutiple++
		sum += divisor
	}
	if mutiple > maxUint32 {
		mutiple = maxUint32
	}
	if signCount == 1 {
		mutiple = -mutiple
	}
	return mutiple
}

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
