package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, tmpSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if tmpSum > 0 {
			tmpSum += nums[i]
		} else {
			tmpSum = nums[i]
		}
		if tmpSum > maxSum {
			maxSum = tmpSum
		}
	}
	return maxSum
}
