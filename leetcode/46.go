package main

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	var res [][]int
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		r := permute(append(nums[:i], nums[i+1:]...))
		for j := range r {
			//fmt.Println("nums[i]" + strconv.Itoa(nums[i]) + strconv.Itoa(t))
			r[j] = append(r[j], t)
			res = append(res, r[j])
		}
	}
	return res
}
