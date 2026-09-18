func isPalindrome(s string) bool {
	// 先頭、末尾それぞれたぐっていって、クロスしたら終わり
	start := 0
	end := len(s)-1
	for {
		// startが endよりも小さい限りloopする
		if start >= end {
			break
		}
		if !unicode.IsLetter(rune(s[start])) && !unicode.IsDigit(rune(s[start])) {
			start++
			continue
		}
		if !unicode.IsLetter(rune(s[end])) && !unicode.IsDigit(rune(s[end])) {
			end--
			continue
		}
		if unicode.ToLower(rune(s[start])) != unicode.ToLower(rune(s[end])) {
			return false
		}
		start++
		end--
	}
	return true
}
