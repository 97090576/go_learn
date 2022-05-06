package main

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	if sameRoot(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func sameRoot(A *TreeNode, B *TreeNode) bool {
	if A == nil || A.Val != B.Val {
		return false
	}
	if A.Val == B.Val {
		if B.Left != nil {
			if !sameRoot(A.Left, B.Left) {
				return false
			}
		}
		if B.Right != nil {
			if !sameRoot(A.Right, B.Right) {
				return false
			}
		}
	}
	return true
}
