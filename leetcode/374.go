package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pick = 6

func guess(num int) int {
	if pick < num {
		return -1
	}
	if pick == num {
		return 0
	}
	return 1
}

func guessNumber(n int) int {
	l := 0
	r := n
	for l < r {
		m := l + (r-l)/2
		if guess(m) == 0 {
			return m
		} else if guess(m) == 1 {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
