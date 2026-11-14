func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures)) // 全て0で、ここに答えが返ってくる
	stack := []int{} // temperaturesのindexを入れる。ここは待合室で自分よりも暖かい気温が入ってくるまで待機する。
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i]  {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		
		stack = append(stack, i)
	}
	return res
}
