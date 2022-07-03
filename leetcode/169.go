package main

func majorityElement(nums []int) int {
	var i, cnt int
	for _, v := range nums {
		if cnt == 0 {
			i = v
		}
		if i == v {
			cnt++
		} else {
			cnt--
		}
	}
	return i
}