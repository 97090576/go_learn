package main

/*
知识点：
1. 队列的使用
2. 二维数组的创建

*/

import (
	"github.com/Workiva/go-datastructures/queue"
)

func dfs(image [][]int, x int, y int, newColor int, oldColor int) {
	if validIndex(x, len(image)) && validIndex(y, len(image[0])) && image[x][y] == oldColor {
		image[x][y] = newColor
		for _, v := range dirs {
			dfs(image, x+v.x, y+v.y, newColor, oldColor)
		}
	}
}

func floodFillDfs(image [][]int, sr int, sc int, newColor int) [][]int {
	dfs(image, sr, sc, newColor, image[sr][sc])
	return image
}

func validIndex(x int, length int) bool {
	return x >= 0 && x < length
}

func floodFillBfs(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	flag := make([][]bool, m)

	i := image[sr][sc]

	for i := range flag {
		flag[i] = make([]bool, n)
	}
	q := queue.New(int64(m * n))
	_ = q.Put(pair{sr, sc})
	flag[sr][sc] = true
	count := 1
	for !q.Empty() {
		get, _ := q.Get(int64(count))
		count = 0
		for _, v := range get {
			p := v.(pair)
			for _, vd := range dirs {
				if validIndex(p.x+vd.x, m) && validIndex(p.y+vd.y, n) && !flag[p.x+vd.x][p.y+vd.y] && image[p.x+vd.x][p.y+vd.y] == i {
					flag[p.x+vd.x][p.y+vd.y] = true
					_ = q.Put(pair{p.x + vd.x, p.y + vd.y})
					count++
				}
			}
			image[p.x][p.y] = newColor
		}
	}
	return image
}
