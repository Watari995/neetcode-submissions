func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices[1:] {
		profit := p - minPrice
		if p < minPrice {
			minPrice = p
		}
		if profit > maxProfit {
			maxProfit = profit
		}
		
	}
	return maxProfit
}
