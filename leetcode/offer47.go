package main

func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([]int, n)
	for i, v := range grid[0] {
		if i == 0 {
			dp[i] = v
		} else {
			dp[i] = dp[i-1] + v
		}
	}
	for i := 1; i < m; i++ {
		tmp := make([]int, n)
		tmp[0] = dp[0] + grid[i][0]
		for j := 1; j < n; j++ {
			tmp[j] = max(tmp[j-1], dp[j]) + grid[i][j]
		}
		dp = tmp
	}
	return dp[n-1]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
