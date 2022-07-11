package main

import "fmt"

func d1() {
	a := []int{1, 2, 4}
	n := 0
	al := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	m := make(map[byte]int)
	for i, v := range al {
		m[v] = i
	}
	fmt.Scan(&n)
	for n > 0 {
		n--
		s := ""
		fmt.Scan(&s)
		bytes := []byte(s)
		for i, v := range bytes {
			if i > len(a) {
				a = append(a, a[i-1]+a[i-2]+a[i-3])
			}
			bytes[i] = al[(m[v]+a[i])%26]
		}
		fmt.Println(string(bytes))
	}
}
