package tree

// 对称的二叉树
func Question28() string {
	tree := []int{1, 2, 2, 3, 4, 4, 3}
	IsSymmetrical(Ints2TreeNode(tree))
	return "done!"
}

func IsSymmetrical(pRoot *TreeNode) bool {
	return isSame(pRoot, pRoot)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val &&
		isSame(root1.Left, root2.Right) &&
		isSame(root1.Right, root2.Left)
}
