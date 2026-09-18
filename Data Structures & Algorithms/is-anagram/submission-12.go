func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, v := range s {
		count[v]++
	}

	// check t with count map
	for _, v := range t {
		_, ok := count[v]
		if !ok || count[v] == 0 {
			return false
		}
		count[v]--
	}
	return true
}
