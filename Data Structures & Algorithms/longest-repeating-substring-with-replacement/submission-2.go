func characterReplacement(s string, k int) int {
	count := [26]int{}
	l, maxFreq, res := 0, 0, 0
	for r := 0; r < len(s); r++ {
		count[s[r]-'A']++
		maxFreq = max(maxFreq, count[s[r]-'A'])

		// maxFreq以外の文字の長さがkよりも大きい時にlを減らしていく
		for (r-l+1) - maxFreq > k {
			count[s[l]-'A']--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
