package main

import (
	"fmt"

	//"myLeetCode_Go/firstweek"
	//"time"
	"myLeetCode_Go/firstweek"
	//"myLeetCode_Go/temp"
)

func main() {
	// 2Sum
	/*	start := time.Now().UnixNano()
		fmt.Println(start)

		var nums = []int{2, 7, 3, 6, 11, 15}
		// var nums = []int{2}

		target := 9
		fmt.Println(firstweek.TwoSumByMap(nums, target))
		end := time.Now().UnixNano()

		fmt.Println(end)

		fmt.Println(end - start)*/

	//3Sum
	/*var num = firstweek.IntSlice{-1, -1, -1, 0, 2, 1, -2, 2, 4, 4}
	//var res [][]int
	fmt.Println(firstweek.ThreeSum(num))

	test := [][]int{{1, 1, 1}, {2, 2, 2}}
	fmt.Println(test)*/

	//temp node
	//初始化一个头结点

	//失败了 尽量不要用go来写链表，用c++更好
	var head firstweek.ListNode
	var n1 firstweek.ListNode
	var n2 firstweek.ListNode
	head.Val = 6
	n1.Val = 4
	n2.Val = 7
	head.Next = &n1
	n1.Next = &n2
	n2.Next = nil

	var head2 firstweek.ListNode
	var m1 firstweek.ListNode
	var m2 firstweek.ListNode
	head2.Val = 5
	m1.Val = 6
	m2.Val = 8
	head2.Next = &m1
	m1.Next = &m2
	m2.Next = nil

	res := firstweek.AddTwoNumbers(&head, &head2)

	fmt.Println(res == nil)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}
