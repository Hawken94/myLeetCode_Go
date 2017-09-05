package main

import (
	"fmt"
	"myLeetCode_Go/forthweek"
	"myLeetCode_Go/sixthweek"
	"myLeetCode_Go/thirdweek"
)

//"myLeetCode_Go/firstweek"
//"time"

//"myLeetCode_Go/temp"

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

	//Add Two Numbers

	//失败了 尽量不要用go来写链表，用c++更好
	/* 	var head firstweek.ListNode
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
	*/

	//Longest Substring Without Repeating Characters
	/* str := "pwwkew"
	fmt.Println(secondweek.LengthOfLongestSubstrings(str))

	//Median of Two Sorted Arrays
	var num1 = []int{1, 4}
	var num2 = []int{2, 3, 4}
	res := secondweek.FindMedianSortedArrays(num1, num2)
	fmt.Println(res)
	*/
	// Longest Palindromic Substring
	fmt.Println(thirdweek.LongestPalindrome("caad"))

	// Reverse Integer
	fmt.Println(forthweek.ReverseWithOverflow(2147483647))

	// String to Integer (atoi)
	fmt.Println(forthweek.MyAtoi("9223372036854775809"))
	// fmt.Println(forthweek.IsNum(' '))

	// Palindrome Number
	fmt.Println(forthweek.IsPalindrome(123))

	// ZigZag Conversion
	fmt.Println(sixthweek.Convert("PAYPALISHIRING", 3))
	// test
}
