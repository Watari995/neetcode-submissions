func trap(height []int) int {
	left, right := 0, len(height) -1
	leftMax, rightMax := height[0], height[len(height)-1]
	total := 0

	for left < right {
		if leftMax <= rightMax {
			left++
			leftMax = max(leftMax, height[left])
			total += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			total += rightMax - height[right]
		}
	}
	return total
}
