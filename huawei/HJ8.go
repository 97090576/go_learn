package main

import (
	"fmt"
	"sort"
)

func mergeTable() {
	var n int
	fmt.Scan(&n)
	m := make(map[int]int, n)
	var k, v int
	for n > 0 {
		n--
		fmt.Scan(&k, &v)
		m[k] += v
	}
	var slice []int
	for i := range m {
		slice = append(slice, i)
	}
	sort.Ints(slice)
	for _, v := range slice {
		fmt.Println(v, m[v])
	}
}
