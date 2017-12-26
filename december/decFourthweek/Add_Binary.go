package decFourthweek

import (
	"strconv"
	"strings"
)

// date:2017-12-24

// 题目:dd Binary
// 给定两个二进制字符串，返回它们的和（也是一个二进制字符串）。
/*
Given two binary strings, return their sum (also a binary string).
For example,
a = "11"
b = "1"
Return "100".
*/

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	c := 0 // 进位标志
	sum := 0
	temp := make([]int, 0)
	for i >= 0 || j >= 0 {
		sum = c
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		c = sum / 2
		temp = append(temp, sum%2)
	}
	if c > 0 {
		temp = append(temp, c)
	}
	// 反转切片
	length := len(temp)
	result := make([]string, len(temp))
	for i := 0; i < len(temp); i++ {
		result[i] = strconv.Itoa(temp[length-i-1])
	}
	return strings.Join(result, "")
}
