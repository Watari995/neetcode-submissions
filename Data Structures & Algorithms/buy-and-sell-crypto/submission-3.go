func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	maxProf := 0
	minPrice := prices[0]
	for i:= 1;i< len(prices);i++ {
		if minPrice > prices[i-1] {
			minPrice = prices[i-1]
		}
		prof := prices[i] - minPrice
		if maxProf < prof {
			maxProf = prof
		}
	}
	return maxProf
}
