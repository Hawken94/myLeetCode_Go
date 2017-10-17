package main

import (
	"fmt"
	"myLeetCode_Go/august/augForthweek"
	"myLeetCode_Go/august/augSixthweek"
	"myLeetCode_Go/constant/quicksort"
	"myLeetCode_Go/july/julFirstweek"
	"myLeetCode_Go/july/julThirdweek"
	"myLeetCode_Go/october/octSecondweek"
	"myLeetCode_Go/october/octThirdweek"
	"myLeetCode_Go/september/sepEighthweek"
	"sort"
	"time"
)

//"myLeetCode_Go/firstweek"
//"time"

//"myLeetCode_Go/temp"

func main() {
	// 2Sum
	start := time.Now().UnixNano()
	fmt.Println(start)

	var nums = []int{2, 7, 3, 6, 11, 15}
	// var nums = []int{2}

	target := 9
	fmt.Println(julFirstweek.TwoSumByMap(nums, target))
	end := time.Now().UnixNano()

	fmt.Println(end)

	fmt.Println(end - start)

	//3Sum
	var num = julFirstweek.IntSlice{-1, -1, -1, 0, 2, 1, -2, 2, 4, 4}
	fmt.Println(julFirstweek.ThreeSum(num))
	nums1 := []int{3, 1, 4, 2}
	quicksort.QuickSort(nums, 0, len(nums1)-1)
	fmt.Println(nums)

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
	   	m2.Val = 85
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
	fmt.Println(julThirdweek.LongestPalindrome("caad"))

	// Reverse Integer
	fmt.Println(augForthweek.ReverseWithOverflow(2147483647))

	// String to Integer (atoi)
	fmt.Println(augForthweek.MyAtoi("9223372036854775809"))
	// fmt.Println(forthweek.IsNum(' '))

	// Palindrome Number
	fmt.Println(augForthweek.IsPalindrome(123))

	// ZigZag Conversion
	fmt.Println(augSixthweek.Convert("PAYPALISHIRING", 3))

	// Container With Most Water
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(augSixthweek.MaxArea(height))

	// Integer to Roman
	fmt.Println(sepEighthweek.IntToRoman2(1996))

	// Roman to Integer
	fmt.Println(sepEighthweek.RomanToInt("MCMXCVI"))

	// Longest Common Prefix
	strs := []string{"aa", "aa"}
	fmt.Println(sepEighthweek.LongestCommonPrefix(strs))
	// fmt.Println("xhk"[0:1])

	// 3Sum Closest
	nums = []int{1, 1, 1, 0}
	target1 := 100
	fmt.Println(sepEighthweek.ThreeSumClosest(nums, target1))

	//
	num1 := sort.IntSlice{-1, 0, 1, 2, -1, -4}
	fmt.Println(octSecondweek.FourSum(num1, -1))

	s := "(){}"
	// b := []byte(s)
	for _, c := range s {
		fmt.Println("xhk", c == '(')
	}

	fmt.Println(octThirdweek.GenerateParentheses(3))
	fmt.Println("xhk" + "xhk")
}
