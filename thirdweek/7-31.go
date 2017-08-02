package thirdweek

import (
	"math"
)

// Longest Palindromic Substring
// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// 给定一个字符串s，找到s中最长的回文子字符串。假设s的最大长度为1000
/*Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example:
Input: "cbbd"
Output: "bb"*/

// LongestPalindrome 中心扩展法
// 思路：分奇偶处理；找到一个中心，向两边扩展，取长度。直到取得最大长度
func LongestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)   //处理奇数情况
		len2 := expandAroundCenter(s, i, i+1) //处理偶数情况
		len := math.Max(float64(len1), float64(len2))
		if int(len) > end-start {
			start = i - int(len-1)/2
			end = i + int(len)/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}

// 马拉车算法解题
