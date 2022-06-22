package leettree

// 二叉树的最小深度
func Leetcode111() string {
	tree := []int{3, 9, 20, NULL, NULL, 15, 7}
	MinDepth(Ints2TreeNode(tree))
	return "done!"
}

func MinDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	nextLevel := make([]*TreeNode, 0)
	nextLevel = append(nextLevel, root)
	for len(nextLevel) != 0 {
		depth += 1 // 上来先加上一层
		temp := make([]*TreeNode, 0)
		for _, node := range nextLevel {
			if node.Left == nil && node.Right == nil {
				// 这就是找到了
				return depth
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		nextLevel = temp
	}
	return depth
}
