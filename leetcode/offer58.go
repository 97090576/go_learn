package main

func reverseLeftWords(s string, n int) string {
	ss := []byte(s)
	res := ss[n:]
	return string(append(res, ss[:n]...))
}
