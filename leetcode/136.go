package main

func singleNumber(nums []int) (r int) {
	for _, v := range nums {
		r ^= v
	}
	return
}
