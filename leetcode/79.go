package main

func exist(board [][]byte, word string) bool {
	flag := make([][]bool, len(board))
	for i := range flag {
		flag[i] = make([]bool, len(board[0]))
	}
	var f func(i int, j int, index int) bool
	f = func(i int, j int, index int) bool {
		if index == len(word)-1 {
			return true
		}
		flag[i][j] = true
		for _, v := range dirs {
			x := i + v.x
			y := j + v.y
			if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && !flag[x][y] && board[x][y] == word[index+1] && f(x, y, index+1) {
				return true
			}
		}
		flag[i][j] = false
		return false
	}
	for i, s := range board {
		for j, v := range s {
			if v == word[0] && f(i, j, 0) {
				return true
			}
		}
	}
	return false
}
