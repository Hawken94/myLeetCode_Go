package eighthweek

// Given an integer, convert it to a roman numeral.
// Input is guaranteed to be within the range from 1 to 3999
// 给定一个整数,将其转换为罗马数字,输入保证在1到3999之间.

// IntToRoman Integer to Roman
// 思路: 建立两个数组将罗马数字和整数对应起来(其实也可以用map); 去数组里找不超过它的最大的数
func IntToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i] //为了获得重复的罗马数字
			result += strs[i]
		}
	}
	return result
}

// IntToRoman2 更加暴力,但是只能是在具体的范围内,有较大的限制
func IntToRoman2(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[num%1000/100] + X[num%100/10] + I[num%10]
}
