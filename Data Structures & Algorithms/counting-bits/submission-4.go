func countBits(n int) []int {
	out := make([]int, n+1)
	for i:=0; i<=n; i++ {
		count := 0
		x := i
	    for x != 0 {
			count += x & 1
			x >>= 1
		}
		out[i] = count
	}
	return out
}
