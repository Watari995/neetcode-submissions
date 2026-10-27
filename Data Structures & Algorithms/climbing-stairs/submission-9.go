func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	n1 := 1
	n2 := 2
	for i := 3; i < n; i++ {
		sum := n1 + n2
		n1 = n2
		n2 = sum
	}
	return n1 + n2
}
