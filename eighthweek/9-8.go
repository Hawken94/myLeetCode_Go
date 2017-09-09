package eighthweek

// Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target.
// Return the sum of the three integers. You may assume that each input would have exactly one solution
// 给定n个整数的数组S，在S中找到三个整数，使得和最接近给定的数目target。返回三个整数的总和。你可以假设每个输入都只有一个解决方案

import (
	"fmt"
	"math"
	"myLeetCode_Go/quicksort"
)

// ThreeSumClosest 3Sum Closest
// 思路：对数组进行排序；算出第一个和第二个数、最后一个数的和，通过绝对值大小来判断
// result的新值
func ThreeSumClosest(nums []int, target int) int {
	quicksort.QuickSort(nums, 0, len(nums)-1)
	result := nums[0] + nums[1] + nums[len(nums)-1]

	for i := 0; i < len(nums)-2; i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum > target {
				end--
			} else {
				start++
			}
			fmt.Println(math.Abs(float64(sum-target)), math.Abs(float64(result-target)))
			if math.Abs(float64(sum-target)) < math.Abs(float64(result-target)) {
				result = sum
				fmt.Printf("xhk: %v\n", result)
			}
		}
	}

	return result
}
