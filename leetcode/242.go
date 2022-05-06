package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for _, v := range s {
		m[v-'a']++
	}
	for _, v := range t {
		if m[v-'a'] == 0 {
			return false
		}
		m[v-'a']--
	}
	return true
}
