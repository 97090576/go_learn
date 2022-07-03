package main

import (
	"bufio"
	"fmt"
	"os"
)

func split8(s string) {
	if len(s) == 0 {
		return
	}
	if len(s) < 8 {
		bytes := []byte(s)
		for len(bytes) < 8 {
			bytes = append(bytes, '0')
		}
		s = string(bytes)
	}
	fmt.Println(s[:8])
	split8(s[8:])
}

func ddd() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		split8(s)
	}
}
