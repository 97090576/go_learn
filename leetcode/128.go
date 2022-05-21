package main

func longestConsecutive(nums []int) (res int) {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for i := range m {
		if !m[i-1] {
			tmp := 1
			for m[i+1] {
				tmp++
				i++
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return
}