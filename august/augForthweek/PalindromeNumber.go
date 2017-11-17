package augForthweek

// date:2017-08-13

// Palindrome Number
// 题目: 确定一个整数是否是回文。不必考虑额外的空间
// Determine whether an integer is a palindrome. Do this without extra space

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
