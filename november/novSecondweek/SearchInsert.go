package novSecondweek

// date:2017-11-09

// SearchInsert
// 题目：给定一个排序数组，返回目标值的所在位置；如果没找到，则返回它在该数组应该所处的位置。假设没有重复值
// Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
// You may assume no duplicates in the array.
// [1,3,5,7,9]  2 → 1

// 类似SearchForRange,二分查找，根据题目不需要找到与target相等的第一个值
func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	} else if nums[len(nums)-1] < target {
		return len(nums)
	}
	result := searchFirst(nums, target)
	return result
}

func searchFirst(nums []int, target int) int {
	low, high := 0, len(nums)
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			if nums[mid+1] >= target {
				return mid + 1
			} else if nums[mid+1] == target {
				return mid + 1
			}
			low = mid + 1
		} else if nums[mid] > target {
			if nums[mid-1] < target {
				return mid
			} else if nums[mid-1] == target {
				return mid - 1
			}
			high = mid
		} else {
			high = mid
		}
	}

	return low
}

// 已经排除了重复值
func searchInsert2(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
