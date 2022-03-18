package leettree

// 最大二叉树
func Leetcode654() string {
	nums := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(nums)
	return "done!"
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var builder func(int, int) *TreeNode
	builder = func(left int, right int) *TreeNode {
		if left == right {
			return &TreeNode{Val: nums[left]}
		}
		if left > right {
			return nil
		}
		var maxVal = 0
		var index = left
		for idx := left; idx <= right; idx++ {
			if maxVal < nums[idx] {
				index = idx
				maxVal = nums[idx]
			}
		}
		var root *TreeNode = &TreeNode{Val: maxVal}
		// 构建左右子树
		root.Left = builder(left, index-1)
		root.Right = builder(index+1, right)
		return root
	}
	return builder(0, len(nums)-1)
}
