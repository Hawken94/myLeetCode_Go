package sepEighthweek

//Given a roman numeral, convert it to an integer.
//Input is guaranteed to be within the range from 1 to 3999.
// 给定一个罗马数字,将其转换为整数,输入保证在1到3999之间.

// RomanToInt Roman to Integer
// 思路:建立一个map,将罗马数字及对应的值映射到map中; 从末端开始遍历, 遇到CM则 M-C, 否则加法
func RomanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := romanMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] >= romanMap[s[i+1]] {
			result += romanMap[s[i]]
		} else {
			result -= romanMap[s[i]]
		}
	}
	return result
}
