package decFirstweek

// date: 2017-12-02

// 题目:
// 在haystack中返回needle第一次出现的索引，如果不是干草堆的一部分，则返回-1
// Implement strStr().
// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

// haystack="aaaaa" needle="bba"
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			// 巧妙之处!!!
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
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
