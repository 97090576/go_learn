package main

import "fmt"

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	paths := make([]string, N)
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&paths[i])
		if err != nil {
			return
		}
	}
	var T int
	_, _ = fmt.Scan(&T)
	qx := make(map[string]byte, T)
	for i := 0; i < T; i++ {
		var s string
		fmt.Scan(&s)
		qx[s[2:]] = s[0]
	}
	m := search(paths, qx)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
