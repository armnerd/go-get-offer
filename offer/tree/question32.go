package tree

// 从上往下打印二叉树
func Question32() string {
	root := []int{8, 6, 10, NULL, NULL, 2, 1}
	PrintFromTopToBottom(Ints2TreeNode(root))
	return "done!"
}

func PrintFromTopToBottom(pRoot *TreeNode) []int {
	if pRoot == nil {
		return nil
	}
	var list = []int{}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	for len(nodeToBeMachined) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		for _, item := range nodeToBeMachined {
			// 先把本节点的放入
			list = append(list, item.Val)
			// 把下层节点先收着下轮处理
			if item.Left != nil {
				nextLevelNodes = append(nextLevelNodes, item.Left)
			}
			if item.Right != nil {
				nextLevelNodes = append(nextLevelNodes, item.Right)
			}
		}
		nodeToBeMachined = nextLevelNodes
	}
	return list
}
