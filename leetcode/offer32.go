package main

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印
func levelOrder_01(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	var res []int
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
// 这样做有一个风险点是所有的节点都放在一个切片里，如果节点过多，可能会 OOM，改进一下，就是每次分配一个新切片保存下一层的节点，可以缓解 OOM 问题
func levelOrder_02(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := []*TreeNode{root}
	// total 表示已经访问过的节点
	total := 0
	// cnt 代表将要访问的那一层的节点
	cnt := 1

	var res [][]int
	for cnt > 0 {
		t := cnt + total
		cnt = 0
		var one []int
		for i := total; i < t; i++ {
			node := q[i]
			one = append(one, node.Val)
			total++
			if node.Left != nil {
				q = append(q, node.Left)
				cnt++
			}
			if node.Right != nil {
				q = append(q, node.Right)
				cnt++
			}
		}
		res = append(res, one)
	}
	return res
}

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
func levelOrder_03(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	q := []*TreeNode{root}
	flag := false
	for len(q) > 0 {
		var p []*TreeNode
		var t []int
		for i := 0; i < len(q); i++ {
			node := q[i]
			t = append(t, node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
		if flag {
			reverse(t, 0, len(t))
		}
		res = append(res, t)
		flag = !flag
	}
	return res
}
