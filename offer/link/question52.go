package link

// 两个链表的第一个公共结点
func Question52() string {
	common := Ints2List([]int{6, 7})
	list1 := Ints2List([]int{1, 2, 3})
	temp1 := list1
	for {
		if temp1.Next == nil {
			temp1.Next = common
			break
		}
		temp1 = temp1.Next
	}
	list2 := Ints2List([]int{4, 5})
	temp2 := list2
	for {
		if temp2.Next == nil {
			temp2.Next = common
			break
		}
		temp2 = temp2.Next
	}
	res := FindFirstCommonNode(list1, list2)
	ReviewAll(res)
	return "done!"
}

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	var ta *ListNode = pHead1
	var tb *ListNode = pHead2
	for {
		if ta == tb {
			break
		}
		if ta != nil {
			ta = ta.Next
		} else {
			ta = pHead2
		}
		if tb != nil {
			tb = tb.Next
		} else {
			tb = pHead1
		}
	}
	return ta
}
