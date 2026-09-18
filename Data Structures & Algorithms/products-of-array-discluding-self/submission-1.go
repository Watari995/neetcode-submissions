func productExceptSelf(nums []int) []int {
	// prefixとsuffixでそれぞれ作成する
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	// suffixを作る、どんどんマイナスしていって、i+1でかけていく
	suffix[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1] 
	}
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[i] = suffix[i] * prefix[i]
	}
	return result
}
