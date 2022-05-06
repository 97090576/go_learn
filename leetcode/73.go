package main

// 矩阵优化空间复杂度的常见方法，使用矩阵的第一行（列）存储结果

func setZeroes(matrix [][]int) {
	// 第一列是否有 0
	col0 := false
	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
			break
		}
	}
	// 标记各行各列
	for _, v := range matrix {
		for j := 1; j < len(v); j++ {
			if v[j] == 0 {
				v[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 置零，注意反向置零，因为第一行和第一列保存着其它列的状态，需要最后才能清除
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
