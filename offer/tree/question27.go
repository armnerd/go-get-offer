package tree

// 二叉树的镜像
func Question27() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	Mirror(root)
	return "done!"
}

func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	temp := pRoot.Left
	pRoot.Left = pRoot.Right
	pRoot.Right = temp
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}
