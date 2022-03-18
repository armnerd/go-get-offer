package tree

// 从上往下打印二叉树
func Question32() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	PrintFromTopToBottom(root)
	return "done!"
}

var list = []int{}

func PrintFromTopToBottom(pRoot *TreeNode) []int {
	if pRoot == nil {
		return nil
	}
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
