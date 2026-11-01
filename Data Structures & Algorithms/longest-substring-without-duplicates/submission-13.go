func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	maxLen := 0
	l := 0
	for r := 0; r < len(s); r++ {
		for seen[s[r]] {
			delete(seen, s[l])
			l++
		}
		seen[s[r]] = true
		size := r - l + 1
		if size > maxLen {
			maxLen = size
		}
	} 
	return maxLen
}
