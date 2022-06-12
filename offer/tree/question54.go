package tree

// 二叉搜索树的第k个结点
func Question54() string {
	tree := []int{5, 3, 7, 2, 4, 6, 8}
	root := Ints2TreeNode(tree)
	KthNode(root, 3)
	return "done!"
}

var res *TreeNode

func KthNode(pRoot *TreeNode, k int) int {
	inOrder(pRoot, &k)
	if res != nil {
		return res.Val
	} else {
		return -1
	}
}

func inOrder(root *TreeNode, k *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, k)
	*k -= 1
	if *k == 0 {
		res = root
		return
	}
	inOrder(root.Right, k)
}
