func threeSum(nums []int) [][]int {
	// まずはsortをする
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 先頭を固定してそこからtwo pointerをやる、重複チェックは最後
	result := [][]int{}
	for i := 0; i < len(nums) - 2; i++ {
		left := i + 1
		right := len(nums) -1
		// 重複チェックして、同じ値なら次に進めてcontinueする
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for left < right {
		   total := nums[i] + nums[left] + nums[right] // これが0になればOK!
		
		if total == 0 {
			result = append(result, []int{nums[i], nums[left], nums[right]})
			right--
			left++
			// leftもチェックして同じ値である限りleft++をする
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for right > left && nums[right] == nums[right+1] {
				right--
			}
		}
		if total > 0 {
			right--
		}
		if total < 0 {
			left++
		}
	    }
	}
	return result
}
