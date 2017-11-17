package novSecondweek

// date:2017-11-08

// Merge k Sorted Lists
// 题目：把k个已排序好的链表合并成一个，分析它的复杂度
// Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

// ListNode 链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二分查找，其余步骤类似MergeTwoLists
func mergeKLists(lists []*ListNode) *ListNode {
	return partion(lists, 0, len(lists)-1)
}
func partion(lists []*ListNode, l, h int) *ListNode {
	if l == h {
		return lists[l]
	}

	if l < h {
		m := (l + h) / 2
		l1 := partion(lists, l, m)
		l2 := partion(lists, m+1, h)
		return merge(l1, l2)
	}
	return nil

}
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
	l2.Next = merge(l1, l2.Next)
	return l2
}
