package decFourthweek

// date:2017-12-19

import (
	"sort"
)

func firstMissingPositive(nums []int) int {
	var temp sort.IntSlice
	temp = nums
	sort.Sort(temp)
	t := 1
	for i := 0; i < len(temp); i++ {
		if temp[i] == t {
			t++
			continue
		}
	}
	return t
}

// 下标交换
func firstMissingPositiveByIndex(nums []int) int {
	length := len(nums)
	i := 0
	for i < length {
		if nums[i] == i+1 || nums[i] <= 0 || nums[i] > length {
			i++
		} else if nums[nums[i]-1] != nums[i] {
			// 下标为i对应的数应该为num[i]-1
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
	i = 0
	for i < length && nums[i] == i+1 {
		i++
	}
	return i + 1
}
