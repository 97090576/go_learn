package main

import "sort"

func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	n := len(nums) - 2
	for n >= 0 {
		j := nums[n]
		s := nums[n+1:]
		reverse(s, 0, len(s))
		i := sort.Search(len(s), func(i int) bool {
			return s[i] > j
		})
		if i != len(s) {
			nums[n], nums[n+1+i] = nums[n+1+i], nums[n]
			return
		}
		reverse(s, 0, len(s))
		n--
	}
	reverse(nums, 0, len(nums))
}
