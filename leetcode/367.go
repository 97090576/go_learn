package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l := 1
	r := num - 1
	for l < r {
		m := l + (r-l)/2
		if m*m == num {
			return true
		} else if m*m > num {
			r = m
		} else {
			l = m + 1
		}
	}
	return false
}
