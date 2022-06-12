package tree

// 二叉搜索树与双向链表
func Question36() string {
	tree := []int{10, 6, 14, 4, 8, 12, 16}
	// tree = []int{5, 4, NULL, 3, NULL, 2, NULL, 1}
	Convert(Ints2TreeNode(tree))
	return "done!"
}

var result *TreeNode = &TreeNode{}
var pre = result

func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return nil
	}
	InOrder(pRootOfTree)
	result.Right.Left = nil
	return result.Right
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	pre.Right = root
	root.Left = pre
	pre = root
	InOrder(root.Right)
}
