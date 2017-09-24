package firstweek

import (
	"fmt"
)

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

//  给定一个整数数组，返回两个数字的索引，使它们相加到一个特定的目标。 假设每个输入都只有一个解决方案，相同的元素只能使用一次

// Example:
// Given nums = [2, 7, 11, 15], target = 9,
// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1]

// TwoSum 利用两个循环遍历
// 时间复杂度O(n^2)
func TwoSum(nums []int, target int) []int {
	//默认数组从小到大排列
	result := make([]int, 0)

	if len(nums) <= 1 {
		return result
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return result
}

// TwoSumByMap 利用一个map
// O(n)
func TwoSumByMap(nums []int, target int) []int {
	result := make([]int, 2)

	if len(nums) <= 1 {
		return result
	}

	tempMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		fmt.Println("xhk:", tempMap[target-nums[i]])
		if _, ok := tempMap[target-nums[i]]; ok {
			result[1] = i // 存储的顺序倒过来了 精妙
			result[0] = tempMap[target-nums[i]]
			// result = append(result, result[1], result[0])   返回多个结果
			return result //只能返回一个结果
		}
		tempMap[nums[i]] = i // 把切片里面的值和索引对应
	}
	return result
}
