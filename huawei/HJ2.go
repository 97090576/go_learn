package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func counts(s string, c byte) (res int) {
	for _, v := range []byte(s) {
		if v == c {
			res++
		}
	}
	return
}

func dd() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s1 := input.Text()
		s1 = strings.ToLower(s1)
		input.Scan()
		s2 := input.Text()
		s2 = strings.ToLower(s2)
		fmt.Println(counts(s1, []byte(s2)[0]))
	}
}
