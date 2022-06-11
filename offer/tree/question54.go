package tree

// 二叉搜索树的第k个结点
func Question54() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	KthNode(root, 3)
	return "done!"
}

var res *TreeNode

func KthNode(pRoot *TreeNode, k int) *TreeNode {
	inOrder(pRoot, &k)
	if res != nil {
		res.Left = nil
		res.Right = nil
	}
	return res
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
