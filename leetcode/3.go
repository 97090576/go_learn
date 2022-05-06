package main

func lengthOfLongestSubstring(s string) int {
	// m 记录 每个字符最后一次出现的位置
	m := make(map[byte]int, len(s))
	// index 记录最后一个重复字符倒数第二个出现的下一个位置
	index := 0
	res := 0
	for i, v := range []byte(s) {
		if m[v] > index {
			index = m[v]
		}
		if i-index+1 > res {
			res = i - index + 1
		}
		m[v] = i + 1
	}
	return res
}
