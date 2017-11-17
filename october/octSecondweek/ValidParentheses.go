package octSecondweek

import (
	"fmt"
	"myLeetCode_Go/utils/stack"
)

// date:unknown

// Valid Parentheses
// 题目:给定一个只包含字符'（'，'）'，'{'，'}'，'['和']'的字符串，确定输入字符串是否有效;括号必须以正确的顺序关闭，“（）”和
// “（）[] {}”全部有效，但“（]”和“（[）]”不是
// Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

// IsValid 思路：由于是判断字符串是否对称，可以栈来设计，入栈和出栈来判断。很巧妙。
func IsValid(s string) bool {
	var stack stack.Stack
	for _, c := range s {
		if c == '(' {
			stack.Push(')')
		} else if c == '{' {
			stack.Push('}')
		} else if c == '[' {
			stack.Push(']')
		} else if stack.IsEmpty() || stack.Pop() != c {
			return false
		}
	}
	fmt.Println(stack)
	return stack.IsEmpty()
}
