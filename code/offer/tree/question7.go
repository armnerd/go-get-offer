package tree

// 重建二叉树
func Question7() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	reConstructBinaryTree(pre, vin)
	return "done!"
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
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
	root_node.Left = reConstructBinaryTree(pre[1:1+index], vin[:index])
	root_node.Right = reConstructBinaryTree(pre[1+index:], vin[index+1:])

	return root_node
}
