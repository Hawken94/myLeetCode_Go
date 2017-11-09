// Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero

// 给定n个整数的数组S，S中有元素a，b，c，使得a+b+c= 0？查找数组中所有独特的三元组，它们的总和为零  (三元组不能重复)

// For example, given array S = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package julFirstweek

import "myLeetCode_Go/utils/quicksort"

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

// ThreeSum 很奇怪，leetcode编译不了，明天先看快速排序算法了
func ThreeSum(nums IntSlice) [][]int {
	// sort.Sort(nums) //对数组进行排序
	quicksort.QuickSort(nums, 0, len(nums)-1)
	//  {-1, -1,-1， 0, 2, 1,2，4，4}
	temp1 := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		min := nums[i]

		for j := i + 1; j < len(nums); j++ {
			if min >= nums[j] {
				min = nums[j]
			}
		}
		temp1 = append(temp1, min)
	}

	var res [][]int
	var temp []int

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			j := i + 1
			k := len(nums) - 1
			sum := 0 - nums[i]

			for j < k {
				if nums[j]+nums[k] == sum {
					temp = append(temp, nums[i], nums[j], nums[k])
					// fmt.Println(temp)
					//temp 最后会变成一个很长的一维数组，不能直接append
					//res = append(res, {nums[i], nums[j], nums[k]})

					for j < k && nums[j] == nums[j+1] {
						j++
					}
					for j < k && nums[k] == nums[k+1] {
						k--
					}
					j++
					k--
				} else if nums[j]+nums[k] < sum {
					j++
				} else {
					k--
				}
			}
		}
	}
	// 哈哈这个算法是我自己想出来的，可以完美输出
	// 将一维切片进行分割后插入二维切片
	ll := len(temp)
	for i := 0; i < ll; i += 3 {
		res = append(res, temp[i:i+3])
	}

	return res
}

func QuickSort2(values []int, left, right int) {
	l, r, temp := left, right, values[left]
	for l < r {
		for l < r && values[r] > temp {
			r--
		}
		if l < r {
			values[l] = values[r]
			l++
		}
		for l < r && values[l] < temp {
			l++
		}
		// 可能已经结束了，赋值时候需要移位
		if l < r {
			values[r] = values[l]
			r--
		}
	}
	values[l] = temp
	QuickSort2(values, left, l-1)
	QuickSort2(values, l+1, right)
}
