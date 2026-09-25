func reverseBits(n int) int {
	res := 0
	for i := 0; i < 32; i++ {
		bit := n & 1 
		n >>= 1

		res <<= 1
		res |= bit
	}
	return res
}
