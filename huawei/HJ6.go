package main

import (
	"fmt"
)

func dddddd() {
	var input int64
	fmt.Scan(&input)

	var i int64

	for i = 2; i*i <= input; i++ {
		//input的质数一定小于等于根号input
		for input%i == 0 {
			fmt.Printf("%d ", i)
			input /= i
		}
	}
	if input >= 2 {
		fmt.Println(input)
	}
}
