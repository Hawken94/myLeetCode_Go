package decFirstweek

// date: 2017-12-02

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
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
