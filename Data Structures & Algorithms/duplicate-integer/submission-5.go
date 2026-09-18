func hasDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
    for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}
