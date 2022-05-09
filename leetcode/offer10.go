package main

func fib(n int) int {
	arr := []int{0, 1}
	if n <= 1 {
		return arr[n]
	}
	return func(n int) int {
		for i := 2; i <= n; i++ {
			arr = append(arr, (arr[i-1]+arr[i-2])%1000000007)
		}
		return arr[n]
	}(n)
}

func numWays(n int) int {
	dp := []int{0, 1, 2}
	if n <= 2 {
		return dp[n]
	}
	return func(n int) int {
		for i := 3; i <= n; i++ {
			dp = append(dp, (dp[i-1]+dp[i-2])%1000000007)
		}
		return dp[n]
	}(n)
}
