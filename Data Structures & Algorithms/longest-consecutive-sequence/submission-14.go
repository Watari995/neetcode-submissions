func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := make(map[int]bool)
	for _, n := range nums {
		numMap[n] = true
	}

	longest := 1

	for n := range numMap {
		if numMap[n-1] {
			continue
		}

		length := 1

		for numMap[n+length] {
			length++
			if length > longest {
				longest = length	
			}
		}
	}

	return longest
}
