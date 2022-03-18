package tree

// 二叉搜索树与双向链表
func Question36() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	Convert(root)
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
