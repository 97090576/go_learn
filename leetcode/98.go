package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	if left != nil {
		for left.Right != nil {
			left = left.Right
		}
	}
	if right != nil {
		for right.Left != nil {
			right = right.Left
		}
	}
	if (left == nil || left.Val < root.Val) && (right == nil || right.Val > root.Val) {
		return isValidBST(root.Left) && isValidBST(root.Right)
	}
	return false
}
