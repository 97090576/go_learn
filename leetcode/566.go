package main

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	if m*n != r*c {
		return mat
	}
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	i := 0
	j := 0
	for _, v := range mat {
		for _, v1 := range v {
			res[i][j] = v1
			j++
			if j == c {
				j = 0
				i++
			}
		}
	}
	return res
}
