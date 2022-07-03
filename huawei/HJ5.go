package main

import (
	"bufio"
	"fmt"
	"os"
)

var m = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'A': 10,
	'B': 11,
	'C': 12,
	'D': 13,
	'E': 14,
	'F': 15,
}

func transform(s string) {
	i := len(s) - 1
	res := 0
	multi := 1
	for i > 1 {
		res += m[s[i]] * multi
		multi *= 16
	}
	fmt.Println(res)
}

func dddd() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		transform(s)
	}
}
