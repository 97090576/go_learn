package main

import (
	"fmt"
)

func main() {
	var value float32 = 0.
	fmt.Scanf("%f", &value)
	run(value)
}

func run(value float32) int {
	v := int(value + 0.5)
	fmt.Printf("%d", v)
	return v
}
