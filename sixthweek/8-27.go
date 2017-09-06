package sixthweek

// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
// Note: You may not slant the container and n is at least 2.
// 给定n个非负整数a1，a2，...，an，其中每个表示坐标点（i，ai）,绘制n条垂直线，使得线i的两个端点在（i，ai）和（i，0）,找到两条线，它们与x轴一起形成一个容器，使得容器含有最多的水.
// 注意：不可以倾斜容器，而且n至少为2
import (
	"math"
)

// MaxArea Container With Most Water
// 思路：从数组的第一个和最后一个开始，下标之差为长，值为宽，宽越小则舍弃，下标值往中间靠拢。图解见官方解析。
func MaxArea(height []int) int {
	max, l, r := float64(0), 0, len(height)-1
	for l < r {
		max = math.Max(float64(max), math.Min(float64(height[l]), float64((height[r])))*float64((r-l)))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return int(max)
}
