package main

func sortColors(nums []int) {
	i := 0
	j := 0
	for i < len(nums) {
		if nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
	if j != len(nums) {
		i = j
		for i < len(nums) {
			if nums[i] == 1 {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			}
			i++
		}
	}
}
