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
		maxLen = max(maxLen, r - l + 1)
	} 
	return maxLen
}
