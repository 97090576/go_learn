package main

func solution(brackets string) bool {
	// 请在在这⾥书写你的代码
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	s := make([]byte, len(brackets))
	for _, v := range []byte(brackets) {
		if v != ' ' {
			b, err := m[v]
			// 右括号
			if err {
				// 栈为空或栈顶左括号与右括号不匹配
				if len(s) == 0 || s[len(s)-1] != b {
					return false
				}
				s = s[:len(s)-1]
			} else {
				s = append(s, v)
			}
		}
	}
	return true
}
