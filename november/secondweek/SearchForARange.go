package secondweek

// Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.

// Your algorithm's runtime complexity must be in the order of O(log n).

// If the target is not found in the array, return [-1, -1].

// For example,
// Given [5, 7, 7, 8, 8, 10] and target value 8,
// return [3, 4].
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
