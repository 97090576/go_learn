package main

func searchMatrix(matrix [][]int, target int) bool {
	// 第一次二分找到 target 可能在的行
	l := 0
	r := len(matrix)
	for l < r {
		m := l + ((r - l + 1) >> 1)
		if matrix[m][0] == target {
			return true
		} else if matrix[m][0] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	// 第二次二分查找是否存在 target
	i := 0
	j := len(matrix[0])
	for i < j {
		m := i + ((j - i) >> 1)
		if matrix[l][m] == target {
			return true
		} else if matrix[l][m] > target {
			j = m
		} else {
			i = m + 1
		}
	}
	return false
}

/* 官方使用 go 的库函数实现
func searchMatrix(matrix [][]int, target int) bool {
    row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
    if row < 0 {
        return false
    }
    col := sort.SearchInts(matrix[row], target)
    return col < len(matrix[row]) && matrix[row][col] == target
}
*/
