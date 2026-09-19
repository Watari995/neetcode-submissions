func countBits(n int) []int {
	out := make([]int, n+1)
	for i:=0;i<=n;i++ {
		count := 0
		x := i
		for x != 0 {
			if x&1 == 1 {
				count++
			}
			x >>= 1
		}
		out[i] = count
	}
	return out
}
