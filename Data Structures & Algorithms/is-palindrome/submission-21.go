func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isLettersOrNumbers(s[l]) {
			l++
		}
		for l < r && !isLettersOrNumbers(s[r]) {
			r--
		}
	if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
		return false
	}
	// move on
	l++
	r--
	}
	return true
}

func isLettersOrNumbers(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' &&  b <= 'Z') || (b >= '0'&& b <= '9')
}