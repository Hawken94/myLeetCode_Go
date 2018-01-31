package fourthweek

// date:2018-1-31
// 题目:插入新的间隔切片 (Insert Interval)
// label:normal

// 给定一个没有重叠的间隔区间的集合,插入一个新的间隔区间使其合并;(有点类似 Merge Intervals)
// Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
// You may assume that the intervals were initially sorted according to their start times.
// Example 1:
// Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].
// Example 2:
// Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].
// This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].

// TODO: leetcode 原题 newInterval 不是指针类型,为了将它赋值为空,特意使用指针类型
func insert(intervals []Interval, newInterval *Interval) []Interval {
	var list []Interval
	// [1,2],[3,5],[6,7],[8,10],[12,16] new:[4,9]
	for _, iList := range intervals {
		if newInterval == nil || iList.End < newInterval.Start {
			list = append(list, iList)
		} else if iList.Start > newInterval.End {
			list = append(list, *newInterval)
			list = append(list, iList)
			newInterval = nil
		} else {
			// 如果有交集
			newInterval.Start = min(iList.Start, newInterval.Start)
			newInterval.End = max(iList.End, newInterval.End)
		}
	}
	if newInterval != nil {
		list = append(list, *newInterval)
	}
	return list
}
