package main

import (
	"fmt"
	"sort"
)

func main() {
	T := 0
	fmt.Scan(&T)
	i := 0
	cnt := 0
	for i <= 45 || i < T {
		var a []int
		i++
		r := i
		s := T / i
		mod := T % i
		l, h := s, s+1
		if mod == 0 {
			r--
			a = append(a, s)
			l--
		}
		for r > 0 {
			if l < 1 {
				break
			}
			a = append(a, l, h)
			l--
			h++
			r -= 2
		}
		if r == 0 {
			if sum(a) == T {
				fmt.Printf("%d=", T)
				sort.Ints(a)
				for i, v := range a {
					if i != 0 {
						fmt.Print("+")
					}
					fmt.Print(v)
				}
				fmt.Print("\n")
				cnt++
			}
		}
	}
	fmt.Printf("Result:%d", cnt)
}

func sum(s []int) (res int) {
	for _, v := range s {
		res += v
	}
	return res
}
