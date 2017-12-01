package decFirstweek

// date:2017-12-01

// Swap Nodes in Pairs
// 题目:交换链表中的结点
// 给定一个链表,交换相邻结点,然后返回交换后的结点.空间复杂度O(n)
/*
Given a linked list, swap every two adjacent nodes and return its head.
For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.
Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.
*/

import (
	"myLeetCode_Go/utils/node"
)

// 画图理解链表
func swapPairs(head *node.ListNode) *node.ListNode {
	var dummy node.ListNode
	dummy.Next = head
	current := &dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next
		first.Next = second.Next

		current.Next = second
		current.Next.Next = first

		current = current.Next.Next
	}
	return dummy.Next
}

func swapPairsByRecursion(head *node.ListNode) *node.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairsByRecursion(head.Next.Next)
	tmp.Next = head
	return tmp
}
