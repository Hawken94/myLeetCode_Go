package decFourthweek

import (
	"strconv"
	"strings"
)

// date:2017-12-20

// 题目: Multiply Strings(大数相乘)
// 给定两个表示为字符串的非负数整数num1和num2，返回num1和num2的乘积
// 注意:不能使用任何内置的BigInteger库或将输入直接转换为整数

/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2.
Note:
The length of both num1 and num2 is < 110.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/

//  `num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`
// 思路:按位相乘,将数据保存到切片中,然后将切片对应位置的值相加,最后再打印切片
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1 := i + j
			p2 := i + j + 1
			sum := int(mul) + pos[p2]

			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}

	// 防止数字开始前有多个0
	str := make([]string, 0)
	for _, v := range pos {
		if v != 0 || len(str) != 0 {
			str = append(str, strconv.Itoa(v))
		}
	}

	s := strings.Join(str, "")
	if len(s) == 0 {
		return "0"
	}
	return s

}
