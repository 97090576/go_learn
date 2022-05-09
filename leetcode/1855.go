package main

import "sort"

// 给你两个 非递增 的整数数组 nums1 和 nums2 ，数组下标均 从 0 开始 计数。
// 二分时间复杂度是 O(n1logn2)
// 使用双指针可以将时间复杂度优化到 O(n1+n2)
// 非递增数组中， nums2[j] >= nums1[i] => nums2[j] >= nums1[i+1], 所以每个数据只需要遍历一遍
func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	for i, v := range nums1 {
		j := sort.Search(len(nums2), func(i int) bool {
			return nums2[i] < v
		})
		if j >= i && j > 0 && j-1-i > res {
			res = j - 1 - i
		}
	}
	return res
}
