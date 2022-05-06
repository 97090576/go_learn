package main

import (
	"sort"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, v := range matrix {
		i := sort.Search(len(v), func(i int) bool { return v[i] >= target })
		if i < len(v) && v[i] == target {
			return true
		}
	}
	return false
}
