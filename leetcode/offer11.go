package main

func minArray(numbers []int) int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		m := l + ((r - l) >> 1)
		if numbers[m] > numbers[r] {
			l = m + 1
		} else if numbers[m] < numbers[r] {
			r = m
		} else {
			l++
		}
	}
	return numbers[l]
}
