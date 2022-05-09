package main

func updateMatrix(mat [][]int) [][]int {
	var q []pair
	for i, v := range mat {
		for j, vv := range v {
			if vv == 0 {
				q = append(q, pair{i, j})
			}
		}
	}
	level := 0
	for len(q) > 0 {
		level--
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, d := range dirs {
				i := v.x + d.x
				j := v.y + d.y
				if i >= 0 && i < len(mat) && j >= 0 && j < len(mat[0]) && mat[i][j] == 1 {
					mat[i][j] = level
					q = append(q, pair{i, j})
				}
			}
		}
	}
	for _, r := range mat {
		for i, v := range r {
			r[i] = -v
		}
	}
	return mat
}
