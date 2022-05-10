package main

func letterCasePermutation(s string) []string {
	n := len(s)
	if n == 0 {
		return []string{s}
	}
	res := letterCasePermutation(s[:n-1])
	if s[n-1] >= '0' && s[n-1] <= '9' {
		for i, v := range res {
			res[i] = string(append([]byte(v), s[n-1]))
		}
	} else {
		var tmp uint8
		if s[n-1] >= 'a' && s[n-1] <= 'z' {
			tmp = s[n-1] + 'A' - 'a'
		} else {
			tmp = s[n-1] + 'a' - 'A'
		}
		for i, v := range res {
			res[i] = string(append([]byte(v), s[n-1]))
			res = append(res, string(append([]byte(v), tmp)))
		}
	}
	return res
}
