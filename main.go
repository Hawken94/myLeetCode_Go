package main

import (
	"fmt"

	//"myLeetCode_Go/firstweek"
	//"time"
	"myLeetCode_Go/temp"
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

	//中午再写  !!!
	var head temp.Node
	var n1 temp.Node
	var n2 temp.Node
	head.Val = 1
	n1.Val = 2
	n2.Val = 3
	*head.Next = *n1.Next
	*n1.Next = *n2.Next

	res := head

	for res.Next != nil {
		fmt.Println(res.Val)
		res = *res.Next
	}
}
