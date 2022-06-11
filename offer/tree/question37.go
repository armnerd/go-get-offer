package tree

import (
	"strconv"
	"strings"
)

// 序列化二叉树
func Question37() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	res := Serialize(root)
	Deserialize(res)
	return "done!"
}

func Serialize(root *TreeNode) string {
	if root == nil {
		return "#,"
	}

	str := strconv.Itoa(root.Val) + ","
	str += Serialize(root.Left)
	str += Serialize(root.Right)
	return str
}

func Deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	str := strings.Split(data, ",")
	index := -1
	var recur func(str []string) *TreeNode
	recur = func(str []string) *TreeNode {
		index++
		if index >= len(str) {
			return nil
		}
		if str[index] == "#" {
			return nil
		}
		value, _ := strconv.Atoi(str[index])
		node := &TreeNode{
			Val:   value,
			Left:  recur(str),
			Right: recur(str),
		}
		return node
	}
	root := recur(str)
	return root
}
