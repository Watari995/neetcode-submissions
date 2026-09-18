func topKFrequent(nums []int, k int) []int {
	// mapを作成して、頻度をチェック、カウントに入れる
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++ // count + 1
	}
	// 数が大きい並び替える
	keys := make([]int, 0, len(numMap))
	for k := range numMap {
		keys = append(keys, k)
	}
	// それをsortする
	sort.Slice(keys, func(i, j int) bool {
		return numMap[keys[i]] > numMap[keys[j]]
	})

	// 先頭の大きいものからkこ取得する
	return keys[:k]
}
