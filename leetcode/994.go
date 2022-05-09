package main

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// 新鲜🍊数量
	cnt := 0
	// 烂🍊
	var q []pair
	for i, v := range grid {
		for j, vv := range v {
			if vv == 1 {
				cnt++
			} else if vv == 2 {
				q = append(q, pair{i, j})
			}
		}
	}
	if cnt == 0 {
		return 0
	}
	// 最长时间
	res := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range dirs {
				i := v.x + vv.x
				j := v.y + vv.y
				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
					grid[i][j] = 2
					q = append(q, pair{i, j})
					cnt--
				}
			}
		}
		res++
	}
	if cnt != 0 {
		res = -1
	}
	return res
}
