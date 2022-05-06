package main

func findRepeatNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		} else {
			m[v]++
		}
	}
	return -1
}
