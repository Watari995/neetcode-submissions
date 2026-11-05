func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]bool)
	l := 0
	maxLen := 0
	for r := 0; r < len(s); r++ {
		for seen[s[r]] {
			delete(seen, s[l])
			l++
		}
		seen[s[r]] = true
		length := r - l + 1
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
