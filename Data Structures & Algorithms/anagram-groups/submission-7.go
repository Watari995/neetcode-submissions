func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)
	for _, str := range strs {
		var cnt [26]int
		for _, s := range str {
			cnt[s-'a']++
		}
		hashMap[cnt] = append(hashMap[cnt], str)
	}
	res := [][]string{}
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}
