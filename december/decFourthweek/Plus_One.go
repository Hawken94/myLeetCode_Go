package decFourthweek

import (
	"math"
)

// date:2017-12-24

// 题目:Plus One
// 给你一个用数组表示的数，求加一之后的结果，结果还是用数组表示。
// 例如:[4,5,6]→[4,5,7] [9,9,9]→[1,0,0,0]
/*
Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
You may assume the integer do not contain any leading zero, except the number 0 itself.
The digits are stored such that the most significant digit is at the head of the list.
*/
func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += c
		c = digits[i] / 10
		digits[i] %= 10
	}
	// 有进位
	if c > 0 {
		temp := make([]int, 0)
		temp = append(temp, 1)
		temp = append(temp, digits...)
		return temp
	}
	math.Sqrt(123)
	return digits
}
