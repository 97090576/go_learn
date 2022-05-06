package main

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	count := 0
	for i, r := range grid {
		for j, v := range r {
			if v == 1 {
				count = 0
				dfs_695(grid, i, j, &count)
				if count > res {
					res = count
				}
			}
		}
	}
	return res
}

func dfs_695(grid [][]int, x int, y int, count *int) {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
		*count++
		grid[x][y] = 0
		for _, v := range dirs {
			dfs_695(grid, x+v.x, y+v.y, count)
		}
	}
}
