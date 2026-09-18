func groupAnagrams(strs []string) [][]string {
	// 単語をkey, 配列をvalueにしてmapを作成する
	strMap := make(map[string][]string)
	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
	    strMap[string(b)] = append(strMap[string(b)], s)
	}
	result := [][]string{}
	// valueを順番に取り出していく
	for _, v := range strMap {
		result = append(result, v)
	}
	return result
}
