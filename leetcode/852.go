package main

func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		} else if arr[m] > arr[m-1] && arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
