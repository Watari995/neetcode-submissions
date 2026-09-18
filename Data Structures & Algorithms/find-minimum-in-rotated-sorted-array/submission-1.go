func findMin(nums []int) int {
	// 切れ目を探す
	// mid と rightを比較して、rightの方が小さかったら切れ目が mid 〜 rightの範囲ある
	// indexで見ていく
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left] 
}
