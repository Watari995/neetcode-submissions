func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, n := range nums {
		numSet[n] = true
	}

	longest := 0 

	// key
	for n := range numSet {
		// ここがスタートラインであることを示す(事前につながるものがない)
		if !numSet[n-1] {
			length := 1

			for numSet[n+length] {
				length++
			}
			
			if length > longest {
				longest = length
			}	
		}
	}
	return longest
}
