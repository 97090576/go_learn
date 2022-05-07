package main

// 生成杨辉三角的前 numRows 行
func generate(numRows int) [][]int {
	res := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		t := make([]int, i+1)
		for j := range t {
			l := 0
			r := 0
			if j != 0 {
				l = res[i-1][j-1]
			}
			if j < i {
				r = res[i-1][j]
			}
			t[j] = l + r
		}
		res = append(res, t)
	}
	return res
}
