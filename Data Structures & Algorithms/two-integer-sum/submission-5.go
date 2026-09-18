func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
    for i, v := range nums {
		// 中身をチェックしてからなければ追加する
		rest := target - v
		if index, ok := numMap[rest]; ok {
			if i < index {
				return []int{i, index}
			}
			return []int{index, i}
		}
		numMap[v] = i
	}
	return []int{}
}
