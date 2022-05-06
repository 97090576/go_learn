package main

import "sort"

// 自写解法，找到大于某个数的最小的数，然后 -1 ，就是小于某个数的最大数，这样两个别扭的地方，一个是需要 -1，另一个是 n 为 1 时需要单独处理
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	l := 1
	r := n
	for l < r {
		m := l + (r-l)/2
		if usedCoin(m) == n {
			return m
		} else if usedCoin(m) > n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}

func usedCoin(n int) int {
	return n * (n + 1) / 2
}

// 官方解法
func arrangeCoins1(n int) int {
	return sort.Search(n, func(k int) bool { k++; return k*(k+1) > 2*n })
}

// 官方解法改写，从左逼近，找到小于某个数的最大数的通用解法
func arrangeCoins2(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r-l+1)/2
		if usedCoin(m) == n {
			return m
		} else if usedCoin(m) > n {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
