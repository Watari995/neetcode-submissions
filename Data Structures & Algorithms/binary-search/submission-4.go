func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		half := (left + right) / 2
		if nums[half] == target {
			return half
		}
		if nums[half] > target { // ターゲットよりも大きかったら左にいく
			right = half - 1
		}
		if nums[half] < target { // ターゲットよりも小さかったら右に行く
			left = half + 1
		}
	}

	return -1
}
