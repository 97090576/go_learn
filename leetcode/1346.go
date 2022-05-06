package main

import (
	"sort"
)

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for j, v := range arr {
		if search1(arr, j+1, len(arr), 2*v) || search1(arr, 0, j, 2*v) {
			return true
		}
	}
	return false
}

func search1(arr []int, l int, r int, target int) bool {
	for l < r {
		m := l + (r-l)/2
		if arr[m] == target {
			return true
		} else if arr[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	return false
}
