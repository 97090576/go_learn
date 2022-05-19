package main

func subsets(nums []int) (res [][]int) {
	var combine []int
	var f func(idx int)
	f = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int(nil), combine...))
			return
		}
		f(idx + 1)
		combine = append(combine, nums[idx])
		f(idx + 1)
		combine = combine[:len(combine)-1]
	}
	f(0)
	return res
}
