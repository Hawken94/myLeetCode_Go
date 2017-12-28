package decFifthweek

import (
	"strings"
)

// date:2017-12-28

// 题目:Length of Last Word
// 返回字符串中最后一个单词的长度
/*
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the
length of last word in the string.
If the last word does not exist, return 0.
Note: A word is defined as a character sequence consists of non-space characters only.
*/
// TODO:leetcode上有很多巧妙的方法
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " {
			count++
		} else {
			break
		}
	}
	return count
}
