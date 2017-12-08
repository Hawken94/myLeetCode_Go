package decFirstweek

import (
	"testing"
)

func TestDecFirstweek(t *testing.T) {

	// t.Errorf("\nxhk %v \n", countAndSay(6))

	nums := []int{1, 1, 2, 2, 3}
	t.Errorf("\n length:%v \n", removeDuplicates(nums))
	// t.Errorf("\n strStr: %v\n", strStr("xhkisgood", "is"))
	t.Errorf("\n strStrAll: %v\n", strStrAll("isxhkisgoodis", "is"))
}
