func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1 := 1
	n2 := 2

	for i:=3; i <= n; i++ {
		// add n2 to n1 
		x := n1
		n1 = n2
		n2 = x + n2
	}
	return n2
}
