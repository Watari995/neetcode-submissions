func missingNumber(nums []int) int {
	missing := len(nums)
	
	for i, num := range nums {
		missing ^= i
		missing ^= num
	}
	return missing
}
