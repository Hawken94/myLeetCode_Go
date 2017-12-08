package decFirstweek

// date:2017-12-05

// 题目: Count And Say
// 计数序列是前五项的整数序列，如下所示
// The count-and-say sequence is the sequence of integers with the first five terms as following:
// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.
/* 就是每一个数字的读法和前一个数字有关 */

import (
	"strconv"
)

// 思路:拼接count和计数的那个数
func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	curr := "1"
	var tmp string

	for i := 1; i < n; i++ {
		tmp = curr
		curr = "" // !!!
		count := 1
		say := tmp[0]

		for j := 1; j < len(tmp); j++ {
			if say != tmp[j] {
				// 不相同的时候拼接,并重置计数器
				curr = curr + strconv.Itoa(count) + strconv.Itoa(int(say-'1'+1))
				count = 1
				say = tmp[j]
			} else {
				count++
			}
		}
		curr = curr + strconv.Itoa(count) + strconv.Itoa(int(say-'1'+1))

	}

	return curr
}
