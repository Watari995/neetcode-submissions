func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices[1:] {
		profit := p - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if p < minPrice {
			minPrice = p
		}
	}
	return maxProfit
}
