package main

func square(num int) int {
	return num * num
}

func sortedSquares(nums []int) []int {
	// 首先 二分查找到 0 或者最小的正数
	l := 0
	r := len(nums)
	var minPos int
	for l < r {
		m := l + (r-l)/2
		if nums[m] == 0 {
			l = m
			r = m
		} else if nums[m] < 0 {
			l = m + 1
		} else {
			r = m
		}
	}
	minPos = l
	maxNeg := l - 1
	// 双指针找平方较小的数加入数组
	res := make([]int, len(nums))
	i := 0
	for minPos < len(nums) && maxNeg >= 0 {
		if square(nums[minPos]) < square(nums[maxNeg]) {
			res[i] = square(nums[minPos])
			minPos++
		} else {
			res[i] = square(nums[maxNeg])
			maxNeg--
		}
		i++
	}
	for minPos < len(nums) {
		res[i] = square(nums[minPos])
		minPos++
		i++
	}
	for maxNeg >= 0 {
		res[i] = square(nums[maxNeg])
		maxNeg--
		i++
	}
	return res
}
