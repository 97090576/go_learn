package main

func firstUniqChar_387(s string) int {
	m := make([]int, 26)
	for i, v := range []byte(s) {
		if m[v-'a'] == 0 {
			// i+1 作用是将 0 值无效化，兼容下标为 0 的字符就是第一个字符的情况
			m[v-'a'] = i + 1
		} else {
			m[v-'a'] = len(s) + 1
		}
	}
	res := len(s) + 1
	for _, v := range m {
		if v > 0 && v < res {
			res = v
		}
	}
	return res%(len(s)+1) - 1
}
