package decFourthweek

import "testing"

func TestCombinationSum(t *testing.T) {
	nums := []int{2, 3, 5, 6, 7}
	target := 7

	t.Errorf("combinationSum : %v \n", combinationSum(nums, target))

	t.Errorf("combinationSum2 : %v \n", combinationSum2(nums, target))

	nums = []int{1, 0, 3, 2}
	t.Errorf("firstMissingPositive :%v \n", firstMissingPositive(nums))

	t.Errorf("firstMissingPositive :%v \n", firstMissingPositiveByIndex(nums))
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	t.Errorf("trap: %v \n", trap(height))

	num1, num2 := "123", "456"
	t.Errorf("multiply :%v %v \n", multiply(num1, num2), len(multiply(num1, num2)))

}
func TestPow(t *testing.T) {
	t.Errorf("myPow :%v \n", myPow(2, -2))
}

func TestPernute(t *testing.T) {
	nums := []int{1, 2, 1}
	t.Errorf("permute :%v \n", permuteUnique(nums))

}

func TestGroupAnagrams(t *testing.T) {
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Errorf("groupAnagrams %v\n", groupAnagrams([]string{""}))
}

func TestMaxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	t.Errorf("maxSubArray :%v \n", maxSubArray(nums))
}

func TestPlusOne(t *testing.T) {
	digits := []int{4, 5, 6}
	t.Errorf("plusOne %v \n", plusOne(digits))
}

func TestAddBinary(t *testing.T) {
	a := "111"
	b := "1"
	t.Errorf("addBinary :%v \n", addBinary(a, b))
}
