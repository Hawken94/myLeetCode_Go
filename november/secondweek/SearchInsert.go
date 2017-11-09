package secondweek

// [1,3,5,7,9]  2
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
