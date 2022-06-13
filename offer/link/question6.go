package link

// 从尾到头打印链表
func Question6() string {
	list := []int{67, 0, 24, 58}
	PrintListFromTailToHead(Ints2List(list))
	return "done"
}

func PrintListFromTailToHead(head *ListNode) []int {
	var tmp = make([]int, 0)
	// 参数判断
	if head == nil {
		return tmp
	}

	// 压入切片
	for {
		tmp = append(tmp, head.Val)
		if head = head.Next; head == nil {
			break
		}
	}
	var res = make([]int, 0, len(tmp))

	// 反转数组
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}
