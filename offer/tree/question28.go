package tree

// 对称的二叉树
func Question28() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	isSymmetrical(root)
	return "done!"
}

func isSymmetrical(pRoot *TreeNode) bool {
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
