func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	max := 0
	for l < r {
		area := (r - l) * min(heights[l], heights[r])
		if area > max {
			max = area
		}
		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}
	return max
}
