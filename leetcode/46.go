package main

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{{nums[0]}}
	}
	res := permute(nums[:n-1])
	for j, v := range res {
		for i := 0; i <= len(v); i++ {
			if i == 0 {
				tmp := []int{nums[n-1]}
				res[j] = append(tmp, v...)
			} else {
				var tmp []int
				tmp = append(tmp, v[:i]...)
				tmp = append(tmp, nums[n-1])
				tmp = append(tmp, v[i:]...)
				res = append(res, tmp)
			}
		}
	}
	return res
}
