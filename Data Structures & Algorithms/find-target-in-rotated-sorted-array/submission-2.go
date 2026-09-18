func search(nums []int, target int) int {
	// 昇順を探してその中でターゲットを探す、それができたら targetの範囲検索をする
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target { return mid } // 要素が一つだった時の対策
		if nums[left] <= nums[mid] { // 左
			// その中にターゲットがいるか
			if nums[left] <= target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
