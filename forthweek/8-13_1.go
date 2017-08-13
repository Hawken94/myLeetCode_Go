package forthweek

// IsPalindrome Palindrome Number
func IsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	// 思路和反转数字一样
	res := 0
	for x > res {
		res = res*10 + x%10
		x /= 10
	}
	// 奇数和偶数的回文
	return (x == res || res/10 == x)
}
