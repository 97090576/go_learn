package main

func buildTree(preorder []int, inorder []int) (res *TreeNode) {
	i := preorder[0]
	res = &TreeNode{Val: i}
	if len(preorder) == 1 {
		return
	}
	index := func(target int) int {
		for i, v := range inorder {
			if v == target {
				return i
			}
		}
		return len(inorder)
	}
	i2 := index(i)
	res.Left = buildTree(preorder[1:i2+1], inorder[:i2])
	res.Right = buildTree(preorder[i2+1:], inorder[i2+1:])
	return
}
