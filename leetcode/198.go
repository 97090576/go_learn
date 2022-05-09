package main

func rob(nums []int) int {
	// dp[i] 表示 偷到 第 i（下标为 i-1） 家房子的最大收入，
	dp := []int{0, nums[0]}
	for i := 2; i <= len(nums); i++ {
		t := dp[i-2] + nums[i-1]
		dp[i] = max(t, dp[i-1])
	}
	return dp[len(nums)]
}
