package link

// 链表中环的入口结点
func Question23() string {
	list := Ints2ListWithCycle([]int{1, 2, 3, 4, 5}, 2)
	EntryNodeOfLoop(list)
	return "done!"
}

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	var fast *ListNode = pHead
	var slow *ListNode = pHead
	for {
		if fast == nil || fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		var empty *ListNode
		return empty
	}
	fast = pHead
	for {
		if fast == slow {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
