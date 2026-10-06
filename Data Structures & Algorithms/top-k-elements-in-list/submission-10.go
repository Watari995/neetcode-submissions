func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, c := range count {
		buckets[c] = append(buckets[c], num)
	}

	res := []int{}
	for i := len(buckets) - 1; i > 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
