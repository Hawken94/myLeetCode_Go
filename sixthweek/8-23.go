package sixthweek

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
