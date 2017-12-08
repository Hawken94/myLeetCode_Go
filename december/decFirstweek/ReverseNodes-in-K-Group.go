package decFirstweek

// date:2017-12-01

// 题目:
// 给定一个链表，一次颠倒链表k的节点并返回其修改列表
// k是一个正整数，小于或等于链表的长度。如果节点的数量不是k的倍数，那末最后的剩余节点应该保持原样
// 您不能更改节点中的值，只有节点本身可能会更改
// 只有恒定的内存是允许的

/*
1.Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
2.k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a
multiple of k then left-out nodes in the end should remain as it is.
2.You may not alter the values in the nodes, only nodes itself may be changed.
3.Only constant memory is allowed

For example,
Given this linked list: 1->2->3->4->5
For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5
*/

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
