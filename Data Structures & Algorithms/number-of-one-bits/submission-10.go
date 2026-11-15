func hammingWeight(n int) int {
	count := 0

	for n > 0 {
		count += n & 1
		n >>= 1 // bitを右にずらしていく
	}
	return count
}
