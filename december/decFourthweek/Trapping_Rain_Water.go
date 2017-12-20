package decFourthweek

// date:2017-12-20

// 题目: Trapping Rain Water
// 给定n个非负整数来表示每个柱的宽度为1的高程图，计算在下雨之后它能够捕获多少水(这个要去官方页面看示意图)

// Given n non-negative integers representing an elevation map where the width of each bar is 1,
// compute how much water it is able to trap after raining.
// For example,
// Given [0,1,0,2,1,0,1,3,2,1,2,1], return 6.

// 思路:用方块与方块之间组成的总面积减去每个方块组成的面积(具体看图)
func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}

	l, r := 0, length-1
	lv, rv := height[l], height[r]
	total, rainTotal := lv+rv, lv+rv
	for l < r {
		if height[l] < height[r] {
			l++
			total += height[l]
			lv = max(lv, height[l])
			rainTotal += lv // lv和height[l]的差值就是积水的面积
		} else {
			r--
			total += height[r]
			rv = max(rv, height[r])
			rainTotal += rv
		}
	}
	return rainTotal - total
}
func max(l, r int) int {
	if l < r {
		return r
	}
	return l
}
