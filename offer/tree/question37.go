package tree

import (
	"strconv"
	"strings"
)

// 序列化二叉树
func Question37() string {
	tree := []int{1, 2, 3, NULL, NULL, 6, 7}
	str := Serialize(Ints2TreeNode(tree))
	Deserialize(str)
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
