package tree

// 二叉树中和为某一值的路径(一)
func Question82() string {
	tree := []int{5, 4, 8, 1, 11, NULL, 9, NULL, NULL, 2, 7}
	HasPathSum(Ints2TreeNode(tree), 22)
	return "done!"
}

func HasPathSum(root *TreeNode, sum int) bool {
	// 遍历到了叶子节点但是该路径不符合条件
	if root == nil {
		return false
	}
	// 目标值递减
	sum -= root.Val
	// 左右节点都为空说明到了叶子节点，和也减到了0，说明找到了符合条件的路径
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	// 对左右分支进行 dfs
	return HasPathSum(root.Left, sum) || HasPathSum(root.Right, sum)
}
