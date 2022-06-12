package tree

// 重建二叉树
func Question7() string {
	var pre = []int{1, 2, 4, 7, 3, 5, 6, 8}
	var vin = []int{4, 7, 2, 1, 5, 3, 8, 6}
	ReConstructBinaryTree(pre, vin)
	return "done!"
}

func ReConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}

	// 找到根节点
	root_val := pre[0]
	root_node := &TreeNode{
		Val: root_val,
	}

	// 切开中序
	var index int
	for index = range vin {
		if vin[index] == root_val {
			break
		}
	}

	// 大树拆小树递归处理
	root_node.Left = ReConstructBinaryTree(pre[1:1+index], vin[:index])
	root_node.Right = ReConstructBinaryTree(pre[1+index:], vin[index+1:])

	return root_node
}
