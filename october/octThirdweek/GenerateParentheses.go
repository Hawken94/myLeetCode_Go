package octThirdweek

// date:unknown

// Generate Parentheses
// 题目:给定n对括号，编写一个函数来生成正确括号的所有组合
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses

// GenerateParentheses 思路:还是看LeetCode的评论吧
func GenerateParentheses(n int) []string {
	var res []string
	helper(n, n, "", &res)
	return res
}

func helper(left, right int, out string, res *[]string) {
	if left < 0 || right < 0 || left > right {
		return
	}
	if left == 0 && right == 0 {
		// fmt.Println(res, ":", out)
		*res = append(*res, out)
		return
	}
	helper(left-1, right, out+"(", res)
	helper(left, right-1, out+")", res)
}
