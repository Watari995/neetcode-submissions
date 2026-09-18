func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
		}
		profit := p - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		
	}
	return maxProfit
}
