package main

import "sort"

func search_offer53(nums []int, target int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] == target
	})
	if i != len(nums) {
		res := 1
		l := i - 1
		r := i + 1
		for l >= 0 && nums[l] == target {
			res++
			l--
		}
		for r < len(nums) && nums[r] == target {
			r++
			res++
		}
		return res
	}
	return 0
}

func missingNumber(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] > i
	})
}
