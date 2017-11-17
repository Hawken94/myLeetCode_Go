package julSecondweek

import "math"

// date:2017-07-30

// Longest Substring Without Repeating Characters
// 题目: 给定一个字符串，找到不重复字符的最长子串的长度。
// Given a string, find the length of the longest substring without repeating characters.

/* Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring. */

// LengthOfLongestSubstrings ... byte是go里面的字节
// 思路：利用map查询，string的[i]值为key，i为value，值相同则进行max比较，取最大
func LengthOfLongestSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	tempMap := make(map[byte]int)
	var j, max float64
	for i := 0; i < len(s); i++ {
		if _, ok := tempMap[s[i]]; ok {
			j = math.Max(float64(j), float64(tempMap[s[i]]+1))
		}
		tempMap[s[i]] = i
		max = math.Max(max, float64(i)-j+1)
	}
	return int(max)
}
