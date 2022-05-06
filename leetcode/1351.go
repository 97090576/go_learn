package main

func countNegatives(grid [][]int) int {
	res := 0
	pos := len(grid[0])
loop:
	for j, v := range grid {
		i := pos
		for i > 0 {
			i--
			if v[i] < 0 {
				res += len(grid) - j
				if i == 0 {
					break loop
				}
			} else {
				pos = i + 1
				break
			}
		}
	}
	return res
}
