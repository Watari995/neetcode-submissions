func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
    for i, v := range nums {
		// 中身をチェックしてからなければ追加する
		rest := target - v
		if index, ok := numMap[rest]; ok {
			sortedNums := []int{index, i}
			if i < index {
				sortedNums = []int{i, index}
			}
			return sortedNums
		}
		numMap[v] = i
	}
	return []int{}
}
