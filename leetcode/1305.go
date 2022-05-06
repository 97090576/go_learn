package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	s1 := treeSlice(root1)
	s2 := treeSlice(root2)
	s3 := make([]int, len(s1)+len(s2))
	i := 0
	j := 0
	k := 0
	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			s3[k] = s1[i]
			i++
		} else {
			s3[k] = s2[j]
			j++
		}
		k++
	}

	// if j == len(s2) {
	//	return append(s3, s1[i:])
	//}
	for i < len(s1) {
		s3[k] = s1[i]
		i++
		k++
	}
	// if i == len(s1) {
	//	return append(s3, s2[j:])
	//}
	for j < len(s2) {
		s3[k] = s2[j]
		j++
		k++
	}
	return s3
}

func treeSlice(root *TreeNode) (s []int) {
	if root == nil {
		return
	}
	l := treeSlice(root.Left)
	s = append(s, l...)
	s = append(s, root.Val)
	r := treeSlice(root.Right)
	s = append(s, r...)
	return
}

// 官方中序遍历
func inorder(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
