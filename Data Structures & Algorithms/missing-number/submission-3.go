func missingNumber(nums []int) int {
	result := len(nums)
	for i, n := range nums {
		result ^= i
		result ^= n
	}
	return result
}
