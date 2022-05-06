package main

func searchRange(nums []int, target int) []int {
	l := 0
	r := len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			l, r = m-1, m+1
			for l >= 0 && nums[l] == target {
				l--
			}
			for r < len(nums) && nums[r] == target {
				r++
			}
			return []int{l + 1, r - 1}
		} else if nums[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return []int{-1, -1}
}
