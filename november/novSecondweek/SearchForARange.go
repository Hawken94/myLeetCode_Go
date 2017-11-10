package novSecondweek

// date:2017-11-09

// SearchForRange
// 题目：给定一个递增数组，找到给定目标值的开始和结束位置，时间复杂度为O(log n);如果没找到，返回[-1,-1]
// Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.
// Your algorithm's runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].
// For example,
// Given [5, 7, 7, 8, 8, 10] and target value 8,
// return [3, 4].

// 思路：采用二分查找，找到target的最开始位置
func searchRange(nums []int, target int) []int {
	start := searchFirstEqual(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	result := make([]int, 0)
	result = append(result, start)
	result = append(result, searchFirstEqual(nums, target+1)-1)
	return result
}

func searchFirstEqual(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
