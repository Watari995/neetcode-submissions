func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, n := range nums {
		countMap[n]++
	}

	bucket := make([][]int, len(nums)+1)
	for num, c := range countMap {
		bucket[c] = append(bucket[c], num)
	}

	res := []int{}
	for i := len(bucket) - 1; i > 0; i-- {
		for _, num := range bucket[i] {
			res = append(res, num)
		}
		if len(res) == k {
			return res
		}
	}
	return res
}
