package fourthweek

import (
	"sort"
)

// date:2018-1-29
// 题目:合并间隔切片 (Merge Intervals)
// label:normal

// 给定间隔的集合，合并所有重叠的间隔。(看例子)
// Given a collection of intervals, merge all overlapping intervals.
// For example,
// Given [1,3],[2,6],[8,10],[15,18],
// return [1,6],[8,10],[15,18].

// Interval 有间隔的
type Interval struct {
	Start int
	End   int
}

// IntervalList 为了实现 sort 接口
type IntervalList []Interval

func merge(intervals IntervalList) []Interval {
	var list IntervalList
	// 对[] Interval 排序
	sort.Sort(intervals)
	// [1,3],[2,6],[8,10],[15,18]
	for _, iList := range intervals {
		// 比较(已经合并的)前一个区间的末元素是否小于后一个区间的首元素(是否有交集)
		if len(list) == 0 || list[len(list)-1].End < iList.Start {
			list = append(list, iList)
		} else {
			// 如果有交集,则取两个区间末元素的最大值
			list[len(list)-1].End = max(list[len(list)-1].End, iList.End)
		}
	}

	return list
}

func (list IntervalList) Len() int {
	return len(list)
}
func (list IntervalList) Less(i, j int) bool {
	if list[i].Start < list[j].Start {
		return true
	}
	return false
}
func (list IntervalList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
