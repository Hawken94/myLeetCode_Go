package decFourthweek

import "sort"

// date:2017-12-23

func permuteUnique(nums []int) [][]int {
	slices := make([][]int, 0)
	// 排序
	var temp sort.IntSlice
	temp = nums
	temp.Sort()
	used := make([]bool, len(nums))
	backtrackUniq(&slices, []int{}, temp, used)
	return slices
}

// 使用一个变量来标记是否使用过了 TODO: debug一下,来看清它的具体流程
func backtrackUniq(slices *[][]int, temp []int, nums []int, used []bool) {
	if len(temp) == len(nums) {
		// slice是引用类型,这里用临时变量防止结果被修改
		tempSlice := make([]int, len(temp))
		copy(tempSlice, temp)
		*slices = append(*slices, tempSlice)
	} else {
		for i := 0; i < len(nums); i++ {
			// TODO: 这个条件不是很懂
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backtrackUniq(slices, temp, nums, used)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
}
