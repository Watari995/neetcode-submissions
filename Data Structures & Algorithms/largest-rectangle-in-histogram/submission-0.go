func largestRectangleArea(heights []int) int {
	stack := []int{}
	maxArea := 0
	n := len(heights)

	for i, h := range heights {
		// 現在のstackのtopよりも低い = 右端が確定した
		for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
			// 右端が確定したバーのインデックスを取り出す
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 幅 = 現在のi - 左側にある自分より低いバーのindex - 1 
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			maxArea = max(maxArea, heights[idx]*width)
		}
		// 右端が未確定なのでstackにつむ
		stack = append(stack, i)
	}


	// ループ終了後も残っているバーは右端=len(heights)で確定
	for len(stack) > 0 {
		idx := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := n
		if len(stack) > 0 {
			width = n - stack[len(stack)-1] - 1
		}
		maxArea = max(maxArea, heights[idx]*width)
	}
	
	return maxArea
}
