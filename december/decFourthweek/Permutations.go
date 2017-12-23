package decFourthweek

// date:2017-12-23

// 题目:Permutations
// 给定一个不重复的数组,返回它的所有的排列组合
/*
Given a collection of distinct numbers, return all possible permutations.
For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

// 排列组合TODO: 传说中的回溯法
func permute(nums []int) [][]int {
	slices := make([][]int, 0)
	backtrackPer(&slices, []int{}, nums)
	return slices
}

func backtrackPer(slices *[][]int, temp []int, nums []int) {
	if len(temp) == len(nums) {
		// slice是引用类型,这里用临时变量防止结果被修改
		tempSlice := make([]int, len(temp))
		copy(tempSlice, temp)
		*slices = append(*slices, tempSlice)
	} else {
		for i := 0; i < len(nums); i++ {
			// temp包含nums[i]就跳过循环
			if isContain(temp, nums[i]) {
				continue
			}
			temp = append(temp, nums[i])
			backtrackPer(slices, temp, nums)
			temp = temp[:len(temp)-1]
		}
	}
}
func isContain(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}
