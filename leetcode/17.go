package main

func letterCombinations(digits string) []string {
	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	for _, v := range []byte(digits) {
		s := m[v]
		if len(res) == 0 {
			bytes := []byte(s)
			for _, v1 := range bytes {
				res = append(res, string(v1))
			}
		} else {
			for _, v := range res {
				for _, v1 := range []byte(s) {
					res = append(res, string(append([]byte(v), v1)))
				}
			}
		}
	}
	return res
}
