package tree

// 把二叉树打印成多行
func Question78() string {
	tree := []int{1, 2, 3, NULL, NULL, 4, 5}
	Print(Ints2TreeNode(tree))
	return "done!"
}

func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	result := [][]int{{pRoot.Val}}
	for len(nodeToBeMachined) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		row := make([]int, 0)
		for _, item := range nodeToBeMachined {
			// 待处理节点始终是从左向右的
			if item.Left != nil {
				nextLevelNodes = append(nextLevelNodes, item.Left)
				row = append(row, item.Left.Val)
			}
			if item.Right != nil {
				nextLevelNodes = append(nextLevelNodes, item.Right)
				row = append(row, item.Right.Val)
			}
		}
		nodeToBeMachined = nextLevelNodes
		if len(row) != 0 {
			result = append(result, row)
		}
	}
	return result
}
