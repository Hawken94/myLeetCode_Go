package decFirstweek

import (
	"strconv"
)

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	curr := "1"
	var tmp string

	for i := 1; i < n; i++ {
		tmp = curr
		curr = ""
		count := 1
		say := tmp[0]

		for j := 1; j < len(tmp); j++ {
			if say != tmp[j] {
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
