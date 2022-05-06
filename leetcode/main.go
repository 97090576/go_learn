package main

import "fmt"

func main() {
	s := []int{1}
	for i := 0; i < len(s); i++ {
		s = append(s, i)
	}
	fmt.Println(s)

}
