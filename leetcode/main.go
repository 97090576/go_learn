package main

import (
	"fmt"
)

func main() {
	r := make([]int, 10, 11)
	r = print(r)
	fmt.Println(r)

}

func print(s []int) []int {
	s = append(s, 1)
	s[9] = 9
	s = append(s, 1)
	fmt.Println(s)
	return s
}
