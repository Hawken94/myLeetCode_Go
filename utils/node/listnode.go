package node

// ListNode 链表定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// GenerateListNode 生成一个链表
func GenerateListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	var list *ListNode
	list.Val = arr[0]
	for i := 1; i < len(arr); i++ {
		list = list.Next
		list.Val = arr[i]
	}

	return list
}
