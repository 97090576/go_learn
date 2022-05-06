package main

func firstUniqChar(s string) byte {
	r := make([]int, 26)
	for _, v := range []byte(s) {
		r[v-'a']++
	}
	for _, v := range []byte(s) {
		if r[v-'a'] == 1 {
			return v
		}
	}
	return ' '
}
