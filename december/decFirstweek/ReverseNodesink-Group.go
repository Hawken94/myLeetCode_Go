package decFirstweek

import (
	"myLeetCode_Go/utils/node"
)

// TODO: 有点复杂,想不通
func reverseKGroup(head *node.ListNode, k int) *node.ListNode {
	if head == nil || head.Next == nil || k < 2 {
		return head
	}

	var dummy node.ListNode
	dummy.Next = head

	tail, prev := &dummy, &dummy
	var temp *node.ListNode
	count := 0
	for {
		count = k
		for count > 0 && tail != nil {
			count--
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		head = prev.Next
		for prev.Next != tail {
			temp = prev.Next
			prev.Next = temp.Next

			temp.Next = tail.Next
			tail.Next = temp
		}
		tail = head
		prev = head
	}
	return dummy.Next

}

func reverseKGroupByRecursion(head *node.ListNode, k int) *node.ListNode {
	curr := head
	count := 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}
	if count == k {
		curr = reverseKGroupByRecursion(curr, k)
		for count > 0 {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
			count--
		}
		head = curr
	}
	return head
}
