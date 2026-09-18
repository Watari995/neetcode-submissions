func singleNumber(nums []int) int {
	result := 0
	for i, _ := range nums {
		result ^= nums[i]
	}	
	return result
}
