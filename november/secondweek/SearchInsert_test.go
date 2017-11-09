package secondweek

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	t.Errorf("xhk: %v \n", searchInsert([]int{1, 4, 4, 4, 4, 4, 4, 4, 4}, 4))
}
