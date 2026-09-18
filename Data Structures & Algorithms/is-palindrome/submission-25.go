func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		// check if it's letters or numbers move right
		for l < r && !isLetters(s[r]) {
			r--
		}
		for l < r && !isLetters(s[l]) {
			l++
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}

		r--
		l++
	}
	return true
}

func isLetters(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}
