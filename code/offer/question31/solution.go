package question31

// 栈的压入、弹出序列
func Run() string {
	pushV := []int{1, 2, 3, 4, 5}
	popV := []int{4, 5, 3, 2, 1}
	IsPopOrder(pushV, popV)
	return "done!"
}

func IsPopOrder(pushV []int, popV []int) bool {
	temp := make([]int, 0)
	i := 0
	j := 0
	all := len(pushV)
	for {
		if i >= all {
			break
		}
		if pushV[i] != popV[j] {
			temp = append(temp, pushV[i])
			i++
		} else {
			// 说明这个元素是放入栈中立马弹出
			i++
			j++
			for {
				// 然后检查popV[j]与栈顶元素是否相等
				if len(temp) == 0 || temp[len(temp)-1] != popV[j] {
					break
				}
				// 如果相等弹出临时栈顶元素
				temp = temp[:len(temp)-1]
				j++
			}
		}
	}
	if len(temp) == 0 {
		return true
	} else {
		return false
	}
}
