package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var f func(target int, pre []int, index int)
	f = func(target int, pre []int, index int) {
		for i := index; i < len(candidates); i++ {
			if target == candidates[i] {
				res = append(res, append(pre, candidates[i]))
				return
			} else if target > candidates[i] {
				f(target-candidates[i], append(pre, candidates[i]), i)
			}
		}
	}
	f(target, []int{}, 0)
	return res
}

/*
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
*/
