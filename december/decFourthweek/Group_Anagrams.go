package decFourthweek

// date:2017-12-24

// 题目:Group Anagrams
// 给定一串字符串，将字母组合在一起
/*
Given an array of strings, group anagrams together.
For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:
[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note: All inputs will be in lower-case.
*/

func groupAnagrams(strs []string) [][]string {
	slices := make([][]string, 0)
	if len(strs) == 0 {
		return nil
	}

	strMap := make(map[string][]string, 0)
	for _, str := range strs {

		data := []byte(str)
		// 排序
		data = BucketSort(data)
		keyStr := string(data)
		strMap[keyStr] = append(strMap[keyStr], string(str))
	}
	for _, strm := range strMap {
		slices = append(slices, strm)
	}
	return slices
}

// BucketSort 桶排序
func BucketSort(data []byte) []byte {

	max := data[0]
	for _, d := range data {
		if max < d {
			max = d
		}
	}
	// 分配桶的个数,此排序范围和data的最大值有关
	sorted := make([]byte, max+1)

	// 将待排序数组里面的数放到与sorted数组相对于的桶里面
	// 待排序的数等于桶排序的索引值，则对应的桶 sorted[data[i]]数值加1
	for i := 0; i < len(data); i++ {
		sorted[data[i]] = sorted[data[i]] + 1
	}
	data = make([]byte, 0)
	for i := 0; i < len(sorted); i++ {
		// 打印出数值不为0的桶的索引值;缺陷:非负整数
		if sorted[i] > 0 {
			// 判断桶里面有多少个数，然后全部打印出来
			for j := 0; j < int(sorted[i]); j++ {
				// fmt.Print(i)
				data = append(data, byte(i))
			}
		}
	}
	return data
}
