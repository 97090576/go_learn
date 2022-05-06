package main

func isValidSudoku(board [][]byte) bool {
	for _, v := range board {
		r := make([]int, 9)
		for _, v1 := range v {
			if v1 != '.' {
				if r[v1-'1'] != 0 {
					return false
				} else {
					r[v1-'1']++
				}
			}
		}
	}
	for i := 0; i < 9; i++ {
		r := make([]int, 9)
		for j := 0; j < 9; j++ {
			v1 := board[j][i]
			if v1 != '.' {
				if r[v1-'1'] != 0 {
					return false
				} else {
					r[v1-'1']++
				}
			}
		}
	}
	x := 0
	for x < 9 {
		y := 0
		for y < 9 {
			r := make([]int, 9)
			for i := x; i < x+3; i++ {
				for j := y; j < y+3; j++ {
					v1 := board[i][j]
					if v1 != '.' {
						if r[v1-'1'] != 0 {
							return false
						} else {
							r[v1-'1']++
						}
					}
				}
			}
			y += 3
		}
		x += 3
	}
	return true
}

/*  本身就是个暴力求解的题，官方在三次遍历的基础上做了时间复杂度的优化，只需要一次遍历
func isValidSudoku(board [][]byte) bool {
    var rows, columns [9][9]int
    var subboxes [3][3][9]int
    for i, row := range board {
        for j, c := range row {
            if c == '.' {
                continue
            }
            index := c - '1'
            rows[i][index]++
            columns[j][index]++
            subboxes[i/3][j/3][index]++
            if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
                return false
            }
        }
    }
    return true
}
*/
