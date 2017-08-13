package forthweek

// Reverse Integer
// Reverse digits of an integer.
// 翻转数字，注意：要考虑溢出情况 （go语言感觉不要考虑溢出）
// Example1: x = 123, return 321
// Example2: x = -123, return -321

/*
// Reverse ...
func Reverse(x int) int {
	str := strconv.Itoa(x)
	str2 := []byte(str)
	n := len(str)
	var tempStr []byte
	for i := 0; i < n; i++ {
		fmt.Println(i)
		tempStr[i] = str2[n-i-1]
	}
	if result, err := strconv.Atoi(string(tempStr)); err != nil {
		return result
	}
	return 0
} */

// Reverse ...
func Reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res
}

// ReverseWithOverflow ...
// 感觉go不需要处理溢出，java要考虑
func ReverseWithOverflow(x int) int {
	res := 0
	for x != 0 {
		temp := x % 10
		newRes := res*10 + temp
		// 如果溢出了，就不会相等，这是作者说的，我也不是很懂，我连溢出都不懂~
		if (newRes-temp)/10 != res {
			return 0
		}
		res = newRes
		x /= 10
	}
	return res
}
