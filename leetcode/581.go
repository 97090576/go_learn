package main

import "math"

func findUnsortedSubarray(nums []int) int {
	l, r := -1, -1
	n := len(nums)
	minNum, maxNum := math.MaxInt, math.MinInt
	for i, num := range nums {
		if num < maxNum {
			r = i
		} else {
			maxNum = num
		}
		if nums[n-i-1] > minNum {
			l = n - i - 1
		} else {
			minNum = nums[n-i-1]
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
