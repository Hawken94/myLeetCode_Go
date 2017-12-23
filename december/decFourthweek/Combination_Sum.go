package decFourthweek

import (
	"sort"
)

// date:2017-12-19

// 题目:
// 给定一个不重复的数组C和一个目标数字T,找到所有元素和为T的组合,元素的次数不限(可重复);所有的数字都为正整数,结果集不能重复.

/*
Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique
combinations in C where the candidate numbers sums to T.
The same repeated number may be chosen from C unlimited number of times.
Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7,
*/

// 思路:回溯法,虽然不是很懂. TODO: 传说中的回溯法
func combinationSum(candidates []int, target int) [][]int {
	slices := make([][]int, 0)
	// 排序
	var temp sort.IntSlice
	temp = candidates
	temp.Sort()
	backtrack(&slices, []int{}, temp, target, 0)
	return slices
}

// {2,3,5,6,7}
func backtrack(slices *[][]int, temp []int, candidates []int, remain, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		// slice是引用类型,这里用临时变量防止结果被修改
		tempSlice := make([]int, len(temp))
		copy(tempSlice, temp)
		*slices = append(*slices, tempSlice)
	} else {
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			backtrack(slices, temp, candidates, remain-candidates[i], i)
			temp = temp[:len(temp)-1]
		}
	}
}
