func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, n := range nums {
		cntMap[n]++
	}

	keys := make([]int, 0, len(cntMap))
	for n := range cntMap {
		keys = append(keys, n)
	} 

	sort.Slice(keys, func(i, j int) bool {
		return cntMap[keys[i]] > cntMap[keys[j]]
	})

	return keys[:k]
}
