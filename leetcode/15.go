package main

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var f func(begin int, target int) [][]int
	f = func(begin int, target int) [][]int {
		end := len(nums) - 1
		var res [][]int
		for begin < end {
			if nums[begin]+nums[end] == target {
				res = append(res, []int{nums[begin], nums[end], -target})
				begin++
				for begin < end && nums[begin] == nums[begin-1] {
					begin++
				}
			} else if nums[begin]+nums[end] < target {
				begin++
			} else {
				end--
			}
		}
		return res
	}
	sort.Ints(nums)
	var res [][]int
	for i, v := range nums {
		if i != 0 && v == nums[i-1] {
			continue
		}
		r := f(i+1, -v)
		if len(r) != 0 {
			res = append(res, r...)
		}
	}
	return res
}
