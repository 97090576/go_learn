package main

import "sort"

// 给定非负整数 c，判断是否存在两个整数 a, b 满足 a² + b² = c
// 设 0 <= a <=b <= c

func judgeSquareSum(c int) bool {
	i := sort.Search(c, func(i int) bool {
		return i*i >= c
	})
	if i*i == c {
		return true
	}
	a := 0
	for a <= i {
		if a*a+i*i == c {
			return true
		} else if a*a+i*i < c {
			a++
		} else {
			i--
		}
	}
	return false
}
