package tree

// 二叉树的深度
func Question55() string {
	var tree = []int{1, 2, 3, 4, 5, NULL, 6, NULL, NULL, 7}
	root := Ints2TreeNode(tree)
	TreeDepth(root)
	return "done!"
}

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		// 说明递归到根部了返回 0
		return 0
	}
	lval := TreeDepth(pRoot.Left)
	rval := TreeDepth(pRoot.Right)
	// 选择左右路径里更长的并算上自己【+1】
	return max(lval, rval) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
