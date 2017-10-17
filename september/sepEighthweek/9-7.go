package sepEighthweek

// LongestCommonPrefix Longest Common Prefix
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
