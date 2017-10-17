package octThirdweek

// GenerateParentheses ...
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
