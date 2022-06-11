package tree

// 二叉搜索树的最近公共祖先
func Question68() string {
	tree := []int{7, 1, 12, 0, 4, 11, 14, NULL, NULL, 3, 5}
	LowestCommonAncestor68(Ints2TreeNode(tree), 1, 12)
	return "done!"
}

func LowestCommonAncestor68(root *TreeNode, p int, q int) int {
	node := core(root, p, q)
	return node.Val
}

// 迭代比较大小
func core(root *TreeNode, p int, q int) *TreeNode {
	if root == nil || p == root.Val || q == root.Val {
		return root
	}
	for root != nil {
		rootVal := root.Val
		if p > rootVal && q > rootVal {
			// 两个节点都在右侧
			root = root.Right
		} else if p < rootVal && q < rootVal {
			// 两个节点都在左侧
			root = root.Left
		} else {
			// 两个节点分别位于左右两侧
			return root
		}
	}
	return nil
}
