package main

// 判断一颗二叉树是不是对称的， 如果一棵二叉树和它的镜像一样，那么它是对称的
// 思路：看到题目的第一印象是求出二叉树的镜像，然后判断两棵树是否相等，这样需要额外的 O(N) 的时间复杂度
// 进一步思考：只需要将右子树翻转，然后判断左右子树是否相等即可
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	root.Right = mirrorTree(root.Right)
	return sameTree(root.Left, root.Right)
}

func sameTree(t1 *TreeNode, t2 *TreeNode) bool {
	return (t1 == nil && t2 == nil) ||
		(t1 != nil && t2 != nil && t1.Val == t2.Val && sameTree(t1.Left, t2.Left) && sameTree(t1.Right, t2.Right))
}

//func mirrorTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	root.Left, root.Right = root.Right, root.Left
//	mirrorTree(root.Left)
//	mirrorTree(root.Right)
//	return root
//}
