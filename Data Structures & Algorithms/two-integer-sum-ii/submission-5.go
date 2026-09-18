func twoSum(numbers []int, target int) []int {
	// 小さければ右に動かす、大きければ左に動かす
	start := 1
	end := len(numbers)
	for {
		if start >= end {
			break
		}

		sum := numbers[start-1] + numbers[end-1]
		
		if sum == target {
			return []int{start, end}
		}

		// 小さければ右
		if sum < target {
			start++
		}

		// 大きければ左
		if sum > target {
			end--
		}
	}
	return []int{}
}
