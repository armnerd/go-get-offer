package tree

// 树的子结构
func Question26() string {
	root1 := []int{8, 8, 7, 9, 2, NULL, NULL, NULL, NULL, 4, 7}
	root2 := []int{8, 9, 2}
	HasSubtree(Ints2TreeNode(root1), Ints2TreeNode(root2))
	return "done!"
}

func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if pRoot1.Val == pRoot2.Val {
		if IsSubtree(pRoot1, pRoot2) {
			return true
		}
	}
	return HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func IsSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	// 先判断 pRoot2
	if pRoot2 == nil {
		return true
	} else if pRoot1 == nil {
		return false
	}
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return IsSubtree(pRoot1.Left, pRoot2.Left) && IsSubtree(pRoot1.Right, pRoot2.Right)
}
