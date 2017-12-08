package decFirstweek

// date:2017-12-02
// 题目:
// 给定一个有序数组，删除重复内容，使每个元素只出现一次，并返回新的长度。 不要为其他数组分配额外的空间，您必须通过在O（1）额外的内存中就地修改
// 输入数组来实现这一点
// Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.
// Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra
//  memory.

// 想输出生成后的切片,但是这个的空间复杂度是O(2)
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
	// fmt.Println(newNums)
	return len(newNums)
}

// 这个方法不必考虑新的切片的具体值
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		// 很巧妙,不相同直接累加
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
