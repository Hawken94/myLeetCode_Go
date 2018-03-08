package secondweek

import (
	"myLeetCode_Go/utils/node"
)

// date:2018-3-8
// 题目:旋转链表(Rotate List)
// lable: 链表 旋转

// 给定一个链表，将链表向右旋转 k 个位置，其中 k 是非负数
// Given a list, rotate the list to the right by k places, where k is non-negative.
// Example:
// Given 1->2->3->4->5->NULL and k = 2,
// return 4->5->1->2->3->NULL.

// 思路:用快慢指针来解，快指针先走k步，然后两个指针一起走，当快指针走到末尾时，慢指针的下一个位置是新的顺序的头结点
func rotateRight(head *node.ListNode, k int) *node.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var dummy node.ListNode
	dummy.Next = head // dummy.Next可以看做头结点

	fast, slow := &dummy, &dummy

	var n int // 链表长度
	for n = 0; fast.Next != nil; n++ {
		fast = fast.Next // 让快指针到达末尾
	}

	for j := n - k%n; j > 0; j-- {
		slow = slow.Next // 让慢指针到达地 n-k%n 个节点
	}

	fast.Next = dummy.Next // 旋转链表
	dummy.Next = slow.Next
	slow.Next = nil

	return dummy.Next

}
