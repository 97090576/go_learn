package main

func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters)
	for l < r {
		m := l + (r-l)/2
		if letters[m] == target {
			m++
			for m < len(letters) && letters[m] == target {
				m++
			}
			l, r = m, m
		} else if letters[m] > target {
			r = m
		} else {
			l = m + 1
		}
	}
	if l == len(letters) {
		l = 0
	}
	return letters[l]
}
