package main

func reverse(nums []int, begin int, end int) {
	l := begin
	r := end - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse(nums, 0, len(nums))
	reverse(nums, 0, k)
	reverse(nums, k, len(nums))
}
