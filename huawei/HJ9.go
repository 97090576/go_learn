package main

import "fmt"

func exitractNum() (res int) {
	var num int
	fmt.Scan(&num)
	m := make(map[int]int)
	for num != 0 {
		i := num % 10
		if m[i] == 0 {
			m[i]++
			res = res*10 + i
		}
	}
	return
}
