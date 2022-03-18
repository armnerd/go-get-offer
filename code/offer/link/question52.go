package link

// 两个链表的第一个公共结点
func Question52() string {
	var common *ListNode = nil
	for _, v := range [2]int{7, 6} {
		node := ListNode{
			Val:  v,
			Next: common,
		}
		common = &node
	}
	var one *ListNode = common
	for _, v := range [3]int{3, 2, 1} {
		node := ListNode{
			Val:  v,
			Next: one,
		}
		one = &node
	}
	var two *ListNode = common
	for _, v := range [2]int{5, 4} {
		node := ListNode{
			Val:  v,
			Next: two,
		}
		two = &node
	}
	res := FindFirstCommonNode(one, two)
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
