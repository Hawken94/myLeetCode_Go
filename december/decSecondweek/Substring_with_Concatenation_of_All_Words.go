package decSecondweek

func findSubstring(s string, words []string) []int {
	strSlice := make([]string, 0)
	for _, word := range words {
		strSlice = append(strSlice, word)
	}
	// str随意组合
	// TODO:用动态规划
}

// 返回所有匹配的下标
func strStrAll(haystack, needle string) []int {
	result := make([]int, 0)
	if len(haystack) < len(needle) {
		return result
	}

	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				if j > 0 {
					j = 0
				}
				result = append(result, i)
				i++
			}
			if i+j == len(haystack) {
				return result
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
