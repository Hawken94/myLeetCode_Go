package sepEighthweek

// date:2017-09-07

//  Longest Common Prefix
// 题目 编写一个函数来查找字符串数组中最长的公共前缀字符串
// Write a function to find the longest common prefix string amongst an array of strings

// LongestCommonPrefix todo 思路:
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	tmpPre := ""
	for i := 1; i < len(strs); i++ {
		if len(pre) == 0 || len(strs[i]) == 0 {
			return ""
		}
		for j := 0; j < len(pre) && j < len(strs[i]); j++ {
			if pre[j] == strs[i][j] {
				tmpPre = pre[:j+1]
			} else {
				tmpPre = pre[:j]
				break
			}
		}
		pre = tmpPre
		if pre == "" {
			break
		}
	}

	return pre
}
