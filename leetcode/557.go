package main

func isSpace(c byte) bool {
	return c == ' '
}

func reverseSlice(s []byte, l int, r int) {
	r = r - 1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseWords(s string) string {
	sSlice := []byte(s)
	l := 0
	r := 0
	for r < len(s) {
		for r < len(s) && !isSpace(sSlice[r]) {
			r++
		}
		reverseSlice(sSlice, l, r)
		r++
		l = r
	}

	return string(sSlice)
}
