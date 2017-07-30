package secondweek

// 7.30写了多道题目
/* There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)). */
// 有两个已经排序好的数组nums1和nums2分别为m和n大小。找到这两个排序数组的中位数。整体运行时间复杂度应为O（log（m + n））
/*
Example 1:
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}

func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// FindMedianSortedArrays ...
//todo 测试通过，但是时间复杂度应该不对
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tempArray = make(IntSlice, 0, 0)
	for i := 0; i < len(nums1); i++ {
		tempArray = append(tempArray, nums1[i])
	}
	fmt.Println(tempArray)
	for i := 0; i < len(nums2); i++ {
		tempArray = append(tempArray, nums2[i])
	}
	fmt.Println(tempArray)

	sort.Sort(tempArray)
	var res float64
	n := len(tempArray)
	if n%2 == 0 {
		res = float64(tempArray[n/2-1]+tempArray[n/2]) / 2
	} else {
		res = float64(tempArray[n/2])
	}
	return res
}
