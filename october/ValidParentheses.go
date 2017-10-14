package october

import (
	"fmt"
	"myLeetCode_Go/constant/stack"
)

// IsValid Valid Parentheses
// 思路：由于是判断字符串是否对称，可以栈来设计，入栈和出栈来判断。很巧妙。
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
