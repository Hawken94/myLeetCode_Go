package main

import (
	"fmt"

	"myLeetCode_Go/firstweek"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	fmt.Println(start)

	var nums = []int{2, 7, 3, 6, 11, 15}
	// var nums = []int{2}

	target := 9
	fmt.Println(firstweek.TwoSumByMap(nums, target))
	end := time.Now().UnixNano()

	fmt.Println(end)

	fmt.Println(end - start)
}
