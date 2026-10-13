func topKFrequent(nums []int, k int) []int {
	cntMap := make(map[int]int)
	for _, n := range nums {
		cntMap[n]++
	}

	// 出現回数 がはいらないので+1をする
	bucket := make([][]int, len(nums)+1)
	for num, cnt := range cntMap {
		bucket[cnt] = append(bucket[cnt], num)
	}

	res := []int{}
	for i := len(bucket) - 1; i >= 0; i-- {
		for _, n := range bucket[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
