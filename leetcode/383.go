package main

// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
func canConstruct(ransomNote string, magazine string) bool {
	m := make([]int, 26)
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		if m[v-'a'] == 0 {
			return false
		}
		m[v-'a']--
	}
	return true
}
