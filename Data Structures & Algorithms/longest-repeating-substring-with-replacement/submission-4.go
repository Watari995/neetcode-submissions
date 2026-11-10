func characterReplacement(s string, k int) int {
	count := [26]int{}
	l, res, maxF := 0, 0, 0
	for r := 0; r < len(s); r++ {
		count[s[r]-'A']++
		if count[s[r]-'A'] > maxF {
			maxF = count[s[r]-'A']
		}
		for (r-l+1) - maxF > k {
			count[s[l]-'A']--
			l++
		}
		if (r-l+1) > res {
			res = r-l+1
		}
	}
	return res
}
