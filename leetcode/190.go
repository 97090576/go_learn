package main

import "fmt"

func reverseBits(n uint32) (rev uint32) {
	fmt.Println(rev)
	for i := 0; i < 32 && n > 0; i++ {
		fmt.Println(n & 1 << (31 - i))
		rev |= n & 1 << (31 - i)
		n >>= 1
		fmt.Println(rev, n)
	}
	return
}
