package main

import (
	"sort"
)

func merge_56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	tmp := intervals[0]
	for _, v := range intervals {
		if v[0] <= tmp[1] {
			tmp[1] = max(v[1], tmp[1])
		} else {
			res = append(res, tmp)
			tmp = v
		}
	}
	res = append(res, tmp)
	return res
}
