package main

func maxSubArray_offer(nums []int) int {
	res, tmp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if tmp > 0 {
			tmp += nums[i]
		} else {
			tmp = nums[i]
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}
