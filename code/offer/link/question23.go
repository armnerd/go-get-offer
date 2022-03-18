package link

// 链表中环的入口结点
func Question23() string {
	var entry = &ListNode{
		Val:  3,
		Next: nil,
	}
	var circle *ListNode
	for k, v := range [2]int{5, 4} {
		var node ListNode
		if k == 0 {
			node = ListNode{
				Val:  v,
				Next: entry,
			}
		} else {
			node = ListNode{
				Val:  v,
				Next: circle,
			}
		}
		circle = &node
	}
	entry.Next = circle
	var start *ListNode
	for k, v := range [2]int{2, 1} {
		var node ListNode
		if k == 0 {
			node = ListNode{
				Val:  v,
				Next: entry,
			}
		} else {
			node = ListNode{
				Val:  v,
				Next: start,
			}
		}
		start = &node
	}
	// ReviewAll(start)
	EntryNodeOfLoop(start)
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
