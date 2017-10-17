package augSixthweek

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//字符串"PAYPALISHIRING"被以给定行的曲折形式写成了下面的样子：(有时候为了方便阅读你可能需要以固定的字体写成这种样子) 注：从上往下再从下往上读再

// convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".

// Convert ZigZag Conversion
func Convert(s string, numRows int) string {
	if len(s) < 0 || numRows <= 1 {
		return s
	}

	var res []byte
	for i := 0; i < len(s) && i < numRows; i++ {
		idx := i
		res = append(res, s[idx])
		for k := 1; idx < len(s); k++ {
			if i == 0 || i == numRows-1 {
				idx += 2*numRows - 2
			} else {
				if (k & 0x1) == 1 {
					idx += 2 * (numRows - 1 - i)
				} else {
					idx += 2 * i
				}
			}

			if idx < len(s) {
				res = append(res, s[idx])
			}
		}
	}
	return string(res)
}
