package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m[v]++
	}
	var res []int
	for _, v := range nums2 {
		if m[v] != 0 {
			m[v]--
			res = append(res, v)
		}
	}
	return res
}
