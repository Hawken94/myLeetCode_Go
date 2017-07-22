// Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero

// 给定n个整数的数组S，S中有元素a，b，c，使得a+b+c= 0？查找数组中所有独特的三元组，它们的总和为零  (三元组不能重复)

// For example, given array S = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package firstweek

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

func ThreeSum(nums IntSlice) [][]int {
	sort.Sort(nums) //对数组进行排序
	//  {-1, -1,-1， 0, 2, 1,2，4，4}

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
					fmt.Println(temp)
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
