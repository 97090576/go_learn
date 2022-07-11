package main

import "fmt"

func ddddddddd() {
	var N int
	fmt.Scan(&N)
	s := make([]int, 501)
	var i int
	for N > 0 {
		fmt.Scan(&i)
		if s[i] == 0 {
			s[i] = 1
		}
	}
	for _, v := range s {
		fmt.Println(v)
	}
}
