func twoSum(numbers []int, target int) []int {
	// key = number, value = false
	hashMap := make(map[int]int)
	for i, n := range numbers {
		hashMap[n] = i+1
	}

	out := make([]int, 2)
	for i, n := range numbers {
		if v, ok := hashMap[target-n]; ok {
			out[0] = i + 1
			out[1] = v
			return out
		}
	}
	return out
}
