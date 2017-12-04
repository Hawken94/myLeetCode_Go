package decFirstweek

import (
	"fmt"
)

// date:2017-12-02

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp := nums[0]
	newNums := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if temp != nums[i] {
			newNums = append(newNums, temp)
			temp = nums[i]
		}
	}
	newNums = append(newNums, temp)
	fmt.Println(newNums)
	return len(newNums)
}

// 这个方法不必考虑新的切片的具体值
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
