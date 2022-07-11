package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	before := make(map[int]int, n)
	after := make(map[int]int, n)
	before[0] = 1
	after[n-1] = 1
	for i := 1; i < n; i++ {
		before[i] = before[i-1] * nums[i-1]
		after[n-1-i] = after[n-i] * nums[n-i]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = before[i] * after[i]
	}
	return res
}
