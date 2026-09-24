func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int)
	var dfs func(n int) int
	dfs = func(n int) int {
		if n >= len(cost) {
			return 0
		}
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		res := cost[n] + min(dfs(n+1), dfs(n+2))
		memo[n] = res
		return res
	}

	return min(dfs(0), dfs(1))
}
