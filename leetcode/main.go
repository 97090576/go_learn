package main

import "fmt"

func main() {
	s := "abc123"
	fmt.Printf("len=%d\n", len(s))
	fmt.Println(s[:4])
}
