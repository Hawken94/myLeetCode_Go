package fourthweek

// date:2018-1-29
// 题目:跳一跳 (Jump Game)
// label:贪心算法

// 给定一个非负整数数组,初始定位在数组的第一个元素,每一个数组的元素代表可以跳的最大距离,判断是否能够到达最后一个点
// (数组的最后一个元素)

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Determine if you are able to reach the last index.
// For example:
// A = [2,3,1,1,4], return true.
// A = [3,2,1,0,4], return false.

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if i > reachable {
			return false
		}
		reachable = max(reachable, i+nums[i])
	}
	return true
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
