func missingNumber(nums []int) int {
	n := len(nums)
	// 0〜nまでの合計を計算するための公式
	sum := n * (n+1) / 2
	for _, v := range nums {
		sum -= v
	}
	return sum
}
