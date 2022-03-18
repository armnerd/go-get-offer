package leettree

// 二叉树的右视图
func Leetcode199() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	rightSideView(root)
	return "done!"
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	nodeToBeMachined := []*TreeNode{root} // 待处理的一层节点
	result = append(result, root.Val)
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
			result = append(result, row[len(row)-1])
		}
	}
	return result
}
