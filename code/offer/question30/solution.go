package question30

var stack []int
var history []int

// 包含min函数的栈
func Run() string {
	Push(-1)
	Push(2)
	Min()
	Top()
	Pop()
	Push(1)
	Top()
	Min()
	return "done!"
}

func Push(node int) {
	stack = append(stack, node)
	if len(history) == 0 {
		history = append(history, node)
	} else {
		top := history[len(history)-1]
		if node < top {
			history = append(history, node)
		} else {
			history = append(history, top)
		}
	}
}

func Pop() {
	stack = stack[:len(stack)-1]
	history = history[:len(history)-1]
}

func Top() int {
	return stack[len(stack)-1]
}

func Min() int {
	return history[len(history)-1]
}
