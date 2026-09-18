func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		// binary search
		middle := (l + r) / 2
		middleNum := nums[middle]
		if middleNum == target {
			return middle
		}

		if middleNum < target {
			l = middle + 1
		} else {
			r = middle - 1
		}
	}
	return -1
}
