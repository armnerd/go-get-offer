package tree

// 在二叉树中找到两个节点的最近公共祖先
func Question86() string {
	tree := []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4}
	LowestCommonAncestor(Ints2TreeNode(tree), 3, 4)
	return "done!"
}

func LowestCommonAncestor(root *TreeNode, p int, q int) int {
	node := Core(root, p, q)
	return node.Val
}

// 递归判断 nil
func Core(root *TreeNode, p int, q int) *TreeNode {
	// 【稍节】没有找到
	if root == nil {
		return nil
	}
	// 找到了其中一个
	if p == root.Val || q == root.Val {
		return root
	}
	left := Core(root.Left, p, q)
	right := Core(root.Right, p, q)
	// 【左右】都没有找到
	if left == nil && right == nil {
		return nil
	}
	// 返回找到的一侧
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// 左右子树上均能找到,那这个节点即是最近公共祖先
	return root
}
