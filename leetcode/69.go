package main

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	l := 0
	r := x
	for l < r {
		m := l + (r-l)/2
		mSquare := m * m
		if mSquare == x {
			return m
		} else if mSquare > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l - 1
}
