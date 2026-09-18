func isValid(s string) bool {
	length := len(s)
	if length % 2 != 0 {
		return false
	}
	last := s[len(s)-1]
	if !isCloseBracket(rune(last)){
		return false
	}

	// for map reference
	bracketMap := map[byte]byte {
		'[' : ']',
		'{' : '}',
		'(' : ')',
 	}

	stack := []rune{}

	for _, v := range s {
		if isCloseBracket(v) {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			value, ok := bracketMap[byte(top)]
			if !ok || value != byte(v) {
				return false
			}
			stack = stack[:len(stack)-1]	
		} else {
		stack = append(stack, v)
		}
	} 
	return len(stack) == 0 
}

func isCloseBracket(r rune) bool {
	return r == ']' || r == ')' || r == '}'
}
