package main

/* 关键点：使用切片实现栈
stack := []int{}
入栈： stack = append(stack, i)
出栈： stack = stack[:len(stack)-1]

*/
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte
	for _, v := range []byte(s) {
		if m[v] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[v] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
