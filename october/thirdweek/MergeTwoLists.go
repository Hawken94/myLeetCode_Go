package thirdweek

// ListNode 链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists 合并两条有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var fakehead ListNode
	l3 := &fakehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			l3.Next = l1
			l1 = l1.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
		}
		l3 = l3.Next
	}
	if l1 != nil {
		l3.Next = l1
	}
	if l2 != nil {
		l3.Next = l2
	}
	return fakehead.Next
}

// MergeTwoLisByRecursive 使用递归
func MergeTwoLisByRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		tmp := l2
		tmp.Next = MergeTwoLisByRecursive(l1, l2.Next)
		return tmp
	} else {
		tmp := l1
		tmp.Next = MergeTwoLisByRecursive(l1.Next, l2)
		return tmp
	}
}
