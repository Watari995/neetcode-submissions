func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0) // 自分より暖かい日が来たら抜けるstackを置いておく	
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			index := stack[len(stack)-1]
			// popする
			stack = stack[:len(stack)-1]
			res[index] = i - index
		}
		stack = append(stack, i)
	}
	return res
}
