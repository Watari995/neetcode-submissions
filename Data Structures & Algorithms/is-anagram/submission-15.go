func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	countMap := make(map[rune]int)

	for _, sValue := range s {
		countMap[sValue]++
	}

	for _, tValue := range t {
		value, ok := countMap[tValue]

		if !ok || value == 0 {
			return false
		}
		countMap[tValue]--
	}
	return true
}
