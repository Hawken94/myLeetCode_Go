package forthweek

// Implement atoi to convert a string to an integer.
// 注意：考虑特殊情况，比如：
//1.该函数首先丢弃尽可能多的空格字符，直到找到第一个非空格字符。 然后，从该字符开始，选择一个可选的初始加号或减号，后跟尽可能多的数字，并将其解释为数值
//2.该字符串可以包含形成整数之后的其他字符，这些字符将被忽略，并且不影响此函数的行为。
//3.如果str中的第一个非空白字符序列不是有效的整数，或者由于str为空或仅包含空格字符，否则不存在此类序列，则不进行转换。 ！！！
//4.如果不能执行有效的转换，则返回零值。如果正确的值超出可表示值的范围，则返回INT_MAX（2147483647）或INT_MIN（-2147483648）

// ！！！只能包含空格
import (
	"math"
	"strings"
)

// MyAtoi String to Integer (atoi)
func MyAtoi(str string) int {
	// 去掉空格，很重要
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	res := 0
	/* var err error
	if res, err = strconv.Atoi(str); err != nil {
		fmt.Println(err)
		return 0
	} */

	var temp []int
	var flag bool
	i := 0
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		flag = true
		i++
	}
	for ; i < len(str); i++ {
		if isNum(str[i]) {
			temp = append(temp, int(str[i]-'0'))
		} else {
			break
		}
	}

	n := len(temp)
	for i := 0; i < n; i++ {
		res += temp[i] * int(math.Pow10(n-1-i))
	}

	if flag {
		res = -res
	}

	// 不完美判断(其他语言的溢出，go好像没有溢出啊),很烦，还是会报错
	if int64(res) > 2147483647 {
		res = 2147483647
	} else if res <= -2147483648 {
		res = -2147483648
	}
	return res
}

// 判断是不是数字，感觉用<'0' >'9'更简洁
func isNum(x byte) bool {
	byteNums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	for _, v := range byteNums {
		if x == v {
			return true
		}
	}

	return false
}
