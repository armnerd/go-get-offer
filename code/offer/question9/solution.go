package question9

import "fmt"

var stack1 []int
var stack2 []int

func Run() string {
	Push(1)
	Push(2)
	Push(3)
	fmt.Println(Pop())
	fmt.Println(Pop())
	Push(4)
	fmt.Println(Pop())
	Push(5)
	fmt.Println(Pop())
	fmt.Println(Pop())
	return "done!"
}

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	// 如果 stack2 空则去 stack1 里倒
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = []int{}
	}

	// 这是 stack2 还没有说明 stack1 也是空的
	if len(stack2) == 0 {
		return -1
	}

	// pop 出 stack2 第一个元素
	index := len(stack2) - 1
	res := stack2[index]
	stack2 = stack2[:index]
	return res
}
