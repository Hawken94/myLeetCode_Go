package secondweek

// date:20189-3-8
// 题目:旋转数组(Rotate Array)
// label: 数组 旋转

/* 将包含 n 个元素的数组向右旋转 k 步。
例如，如果n = 7,k = 3，给定数组[1,2,3,4,5,6,7]，向右旋转后的结果为 [5,6,7,1,2,3,4]。
注意:
尽可能找到更多的解决方案，这里最少有三种不同的方法解决这个问题。
提示:
要求空间复杂度为 O(1)
*/
/* Rotate an array of n elements to the right by k steps.
For example, with n = 7 and k = 3, the array [1,2,3,4,5,6,7] is rotated to [5,6,7,1,2,3,4].
Note:
Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
*/
// 需要 O(n)的额外空间
func rotate(nums []int, k int) {
	length := len(nums)
	result := make([]int, length) // 初始化没给足够的空间
	for i := 0; i < length; i++ {
		result[(i+k)%length] = nums[i]
	}
	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}
}

// 思路:当我们把数组旋转 k 次时,后面的 k 个元素移到了前面,剩下的前面的元素移到后面;先旋转整个数组,然后旋转前 k 个
// 元素,最后旋转 n-k 个元素
func rotate2(nums []int, k int) []int {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	return nums
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}
