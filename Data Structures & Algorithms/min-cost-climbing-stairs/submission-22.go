func minCostClimbingStairs(cost []int) int {
	// keyはindexでvalueはそのindexからのmincostがはいる
	memo := make(map[int]int)
	var minCost func(idx int) int
	minCost = func(idx int) int {
		if idx > len(cost) - 1 {
			return 0
		}
		if value, ok := memo[idx]; ok {
			return value
		}
		memo[idx] = cost[idx] + min(minCost(idx+1), minCost(idx+2))
		return memo[idx]
	}	
	return min(minCost(0), minCost(1))
}
