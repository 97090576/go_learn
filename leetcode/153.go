package main

func findMin(nums []int) int {
	n := len(nums)
	L := nums[0]
	R := nums[n-1]
	// L < R 代表旋转了 n 次
	if L <= R {
		return L
	}
	l := 0
	r := n
	for l < r {
		m := l + (r-l)/2
		if nums[m] >= L {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
