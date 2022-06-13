package link

// 复杂链表的复制
func Question35() string {
	return "done!"
}

func Clone(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	releationship := make(map[*RandomListNode]*RandomListNode)
	sentry := &RandomListNode{}
	pre := sentry
	cur := head
	for cur != nil {
		newOne := &RandomListNode{
			Label: cur.Label,
		}
		releationship[cur] = newOne
		pre.Next = newOne
		pre = pre.Next
		cur = cur.Next
	}
	for origin, newOne := range releationship {
		if origin.Random == nil {
			newOne.Random = nil
		} else {
			newOne.Random = releationship[origin.Random]
		}
	}
	return sentry.Next
}
