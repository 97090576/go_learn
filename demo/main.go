package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 3}, {5, 8}, {2, 6}, {20, 50}, {25, 55}}
	slice := mergeSlice(arr)
	fmt.Println(slice)
}

func mergeSlice(s [][]int) (res [][]int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})
	i := 0
	l, r := 0, 0
	for i < len(s) {
		l = s[i][0]
		r = s[i][1]
		j := i + 1
		for j < len(s) {
			if s[j][0] <= r {
				r = s[j][1]
			} else {
				res = append(res, []int{l, r})
				break
			}
			j++
		}
		i = j
	}
	res = append(res, []int{l, r})
	return
}
