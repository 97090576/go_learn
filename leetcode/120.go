package main

import (
	"math"
	"sort"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := []int{triangle[0][0]}
	for i := 1; i < n; i++ {
		tmp := make([]int, i+1)
		tmp[0] = triangle[i][0] + dp[0]
		for j := 1; j <= i; j++ {
			t := dp[j-1]
			if j != i {
				t = int(math.Min(float64(dp[j]), float64(dp[j-1])))
			}
			tmp[j] = t + triangle[i][j]
		}
		dp = tmp
	}
	sort.Ints(dp)
	return dp[0]
}
