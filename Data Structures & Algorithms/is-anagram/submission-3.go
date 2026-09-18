func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// それぞれのmapを作って数を比較してそこに入れる
	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		// 数を増やす
		countS[ch]++
		// Tも一緒に増やす (iで順番にincrementしている)
		countT[rune(t[i])]++
	}

	// keyとvalueでそれぞれを順番にチェックする
	for k, v := range countS {
		if countT[k] != v {
			return false
		}
	}
	return true
}
