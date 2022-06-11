package tree

func Question55() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
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
