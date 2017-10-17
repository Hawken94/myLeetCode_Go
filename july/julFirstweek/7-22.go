// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// 给定两个非空的链表代表两个非负整数，数字以相反的顺序存储，每个节点包含一个数字，使这两个数字相加并以链表返回。

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

package julFirstweek

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 不懂啊 todo
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyHead ListNode
	p, q, curr := l1, l2, dummyHead

	carry := 0

	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum := carry + x + y
		carry = sum / 10

		var newNode ListNode
		newNode.Val = sum % 10
		curr.Next = &newNode

		fmt.Println(curr.Val)
		curr = *curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		var newNode ListNode
		newNode.Val = carry
		curr.Next = &newNode
	}
	fmt.Println(curr.Next)
	return dummyHead.Next
}
