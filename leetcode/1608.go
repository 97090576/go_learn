package main

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	l, n, r := 1, len(nums), len(nums)
	for l <= r {
		m := l + (r-l)>>1
		if nums[n-m] >= m && (n == m || nums[n-m-1] < m) {
			return m
		} else if nums[n-m] < m {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
