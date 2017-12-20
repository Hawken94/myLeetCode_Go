package decFourthweek

// date:2017-12-19

// 题目: First Missing Positive
// 给定一个未分类的整数数组，找到第一个缺少的正整数(从1开始数)

/*
Given an unsorted integer array, find the first missing positive integer.
For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.
Your algorithm should run in O(n) time and uses constant space.
*/
import (
	"sort"
)

// 通过排序,然后再比较对应位置的值是否正确(排序算法可能不是O(n))
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

// 下标交换,使对应位置的值与对应的下标符合
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
