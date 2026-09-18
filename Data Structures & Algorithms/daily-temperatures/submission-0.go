func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures)) // ゼロ値で初期化をする
	stack := []int{}
	for i, temp := range temperatures {
		// 値がstackに積まれているものより大きい限りずっと回し続けてresultに入れる
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[idx] = i - idx
		}
		stack = append(stack, i) // indexを追加
	}
	return result
}
