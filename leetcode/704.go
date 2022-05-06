package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums)
	m := l + (r-l)/2
	for l < r {
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m
		} else {
			l = m
		}
		m = l + (r-l)/2
	}
	return -1
}
