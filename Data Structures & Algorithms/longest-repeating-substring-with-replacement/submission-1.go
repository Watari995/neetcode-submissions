func characterReplacement(s string, k int) int {
	count := [26]int{}
	l, maxFreq := 0, 0

	for r := 0; r < len(s); r++ {
		count[s[r]-'A']++
		if count[s[r]-'A'] > maxFreq {
			maxFreq = count[s[r]-'A']
		}

		if (r-l+1) - maxFreq > k {
			count[s[l]-'A']--
			l++
		}
	}
	return len(s)-l
}
