package main

import (
	"fmt"
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)
	var count int
	for _, v := range arr1 {
		if !bSearch(v, arr2, float64(d)) {
			count++
			fmt.Println(v)
		}
	}
	return count
}

func bSearch(num int, arr2 []int, d float64) bool {
	l := 0
	r := len(arr2)
	for l < r {
		m := l + (r-l)/2
		if math.Abs(float64(num-arr2[m])) <= d {
			return true
		} else if num < arr2[m] {
			r = m
		} else {
			l = m + 1
		}
	}
	return false
}
