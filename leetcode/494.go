package main

func findTargetSumWays(nums []int, target int) (res int) {
	var f func(arr []int, target int)
	f = func(arr []int, target int) {
		if len(arr) == 0 {
			if target == 0 {
				res++
			}
			return
		}
		f(arr[1:], target+arr[0])
		f(arr[1:], target-arr[0])
	}

	f(nums, target)
	return
}