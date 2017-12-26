package decFourthweek

// date:2017-12-24

// 题目:Maximum Subarray
// 在一个数组中找到连续的子数组（至少包含一个数字），这个数组的总和最大
/*
Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
the contiguous subarray [4,-1,2,1] has the largest sum = 6.
*/
// 思路:记录相邻两个数之和的最大值,然后再记录一段数字的最大值
func maxSubArray(nums []int) int {
	maxSofar := nums[0]
	maxEndHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndHere = max(maxEndHere+nums[i], nums[i])
		maxSofar = max(maxSofar, maxEndHere)
	}
	return maxSofar
}
