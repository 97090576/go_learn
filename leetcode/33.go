package main

// 在旋转有序数组中查找特定值 target 的索引，不存在则返回 -1
func search_33(nums []int, target int) int {
	l := 0
	r := len(nums)
	L := nums[l]
	R := nums[r-1]
	for l < r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} else if (nums[m] > L && (nums[m] < target || (nums[m] > target && target < L))) || (nums[m] < R && nums[m] < target && target < R) {
			l = m + 1
		} else {
			r = m
		}

	}
	return -1
}
