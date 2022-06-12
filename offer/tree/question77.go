package tree

// 按之字形顺序打印二叉树
func Question77() string {
	tree := []int{1, 2, 3, NULL, NULL, 4, 5}
	root := Ints2TreeNode(tree)
	SnakePrint(root)
	return "done!"
}

func SnakePrint(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	result := [][]int{{pRoot.Val}}
	reverse := true // 方向标识
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
		length := len(row)
		if length != 0 {
			if reverse {
				for i := 0; i < length/2; i++ {
					row[length-1-i], row[i] = row[i], row[length-1-i]
				}
			}
			result = append(result, row)
		}
		reverse = !reverse // 转向
	}
	return result
}
