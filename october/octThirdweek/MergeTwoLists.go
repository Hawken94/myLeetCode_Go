package octThirdweek

// ListNode 链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists 思路：以一条链表l2为基准，当其他的链表元素比它大，直接移动了l1，否则移动l2
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

// MergeTwoLisByRecursive 思路：使用递归;如果l1.Val > l2.Val,则l1继续与l2.Next比较,以此类推;反之,类似.
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
	}
	tmp := l1
	tmp.Next = MergeTwoLisByRecursive(l1.Next, l2)
	return tmp

}
