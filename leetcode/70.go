package main

func climbStairs(n int) int {
	dp := []int{1, 1, 2, 3}
	if n <= 3 {
		return dp[n]
	}
	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
