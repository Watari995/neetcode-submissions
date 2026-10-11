func groupAnagrams(strs []string) [][]string {
	countMap := make(map[[26]int][]string)
	for _, s := range strs {
		var cnt [26]int
		for _, c := range s {
			cnt[c-'a']++
		}
		countMap[cnt] = append(countMap[cnt], s)
	}

	res := make([][]string, 0, len(countMap))
	for _, v := range countMap {
		res = append(res, v)
	}
	return res
}
