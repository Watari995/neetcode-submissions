func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	n1 := 1
	n2 := 2

	for i := 3;i <= n;i++ {
		x := n2
		n2 = n1 + n2
		n1 = x
	} 
	return n2
}
