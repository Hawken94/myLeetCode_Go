package october

import (
	"fmt"
	"sort"
)

// FourSum based on 3Sum
// 思路:基于3Sum的想法,排序后先判断前四个,再判断第一个和最后三个;接着是在循环里面判断前三个和最后一个的和与target的比较
func FourSum(num sort.IntSlice, target int) [][]int {
	var ans [][]int
	length := len(num)
	if length < 4 {
		return ans
	}
	sort.Sort(num)
	fmt.Println("xhk:", num)
	for i := 0; i < length-3; i++ {
		if num[i]+num[i+1]+num[i+2]+num[i+3] > target {
			break
		}
		if num[i]+num[length-1]+num[length-2]+num[length-3] < target {
			continue
		}
		if i > 0 && num[i] == num[i-1] { // 去掉重复的
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if num[i]+num[j]+num[j+1]+num[j+2] > target {
				break
			}
			// 每次比较前两个和后两个
			if num[i]+num[j]+num[length-1]+num[length-2] < target {
				continue
			}
			if j > i+1 && num[j] == num[j-1] {
				continue
			}
			low, high := j+1, length-1
			for low < high {
				sum := num[i] + num[j] + num[low] + num[high]
				if sum == target {
					ans = append(ans, []int{num[i], num[j], num[low], num[high]})
					for low < high && num[low] == num[low+1] {
						low++
					}
					for low < high && num[high] == num[high-1] {
						high--
					}
					low++
					high--
				} else if sum < target {
					low++
				} else {
					high--
				}
			}
		}
	}
	return ans
}
