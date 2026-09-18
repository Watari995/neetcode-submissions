func maxArea(heights []int) int {
	// 左右を比較して高い方を優先、低い方はずらす
	maxSize := 0
	left := 0
	right := len(heights) - 1 
	for left < right {
		width := right - left
		height := min(heights[left], heights[right]) // 小さい方を高さにする
		size := width * height
		maxSize = max(maxSize, size)
		if heights[left] > heights[right] {
			right--
		} else {
			left++
		}
	}
	return maxSize
}
