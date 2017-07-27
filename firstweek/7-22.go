// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// 给定两个非空的链表代表两个非负整数，数字以相反的顺序存储，每个节点包含一个数字，使这两个数字相加并以链表返回。

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8

package firstweek

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// }
