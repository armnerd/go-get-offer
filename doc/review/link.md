# 链表

* 从尾到头打印链表
* 反转链表
* 合并两个排序的链表
* 两个链表的第一个公共结点
* 链表中环的入口结点
* 链表中倒数最后k个结点
* 删除链表中重复的结点
* 删除链表的节点

# 6. 从尾到头打印链表

## 题目

输入一个链表，按链表从尾到头的顺序返回一个 ArrayList。

## 思路

遍历链表压入临时切片再反转

## 实现

```go
func printListFromTailToHead(head *ListNode) []int {
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
```

------

# 24. 反转链表

## 题目

给定一个单链表的头结点pHead，长度为n，反转该链表后，返回新链表的表头。

数据范围： n≤1000
要求：空间复杂度 O(1) ，时间复杂度 O(n) 。

如当输入链表{1,2,3}时，
经反转后，原链表变为{3,2,1}，所以对应的输出为{3,2,1}。

## 思路

原地不动改方向

## 实现

```go
func ReverseList(pHead *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode = pHead
	var nex *ListNode
	for {
		if cur == nil {
			break
		}
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
```

------

# 25. 合并两个排序的链表

## 题目

输入两个递增的链表，单个链表的长度为n，合并这两个链表并使新链表中的节点仍然是递增排序的。
数据范围： 0≤n≤10000，−1000≤节点值≤1000
要求：空间复杂度 O(1)，时间复杂度 O(n)

如输入{1,3,5},{2,4,6}时，合并后的链表为{1,2,3,4,5,6}，所以对应的输出为{1,2,3,4,5,6}

或输入{-1,2,4},{1,3,4}时，合并后的链表为{-1,1,2,3,4,4}，所以对应的输出为{-1,1,2,3,4,4}

## 思路

运用游标

## 实现

```go
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	result := ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := &result
	for {
		if pHead1 == nil || pHead2 == nil {
			break
		}
		if pHead1.Val >= pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		}
		cur = cur.Next
	}
	if pHead1 != nil {
		cur.Next = pHead1
	} else if pHead2 != nil {
		cur.Next = pHead2
	}
	return result.Next
}
```

------

# 52. 两个链表的第一个公共结点

## 题目

输入两个无环的单向链表，找出它们的第一个公共结点，如果没有公共节点则返回空。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

数据范围： n≤1000
要求：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

输入：{1,2,3},{4,5},{6,7}
返回值：{6,7}
说明：第一个参数{1,2,3}代表是第一个链表非公共部分，第二个参数{4,5}代表是第二个链表非公共部分，最后的{6,7}表示的是2个链表的公共部分，这3个参数最后在后台会组装成为2个两个无环的单链表，且是有公共节点的

## 思路

* 双指针【假设两个链表长度相等】
* 如何让本来长度不相等的变为相等的？
* 假设链表A长度为a， 链表B的长度为b，此时a != b
* 但是，a+b == b+a
* 因此，可以让a+b作为链表A的新长度，b+a作为链表B的新长度
* a+c+b+c = b+c+a+c 【c是公共部分，长度补偿后依然可以相遇】

## 实现

```go
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
```

------

# 23. 链表中环的入口结点

## 题目

给一个长度为n链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

数据范围： n≤10000
节点值范围：[1,10000]
要求：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

例如，输入{1,2},{3,4,5}时
输入分为2段，第一段是入环前的链表部分，第二段是链表环的部分，后台会根据第二段是否为空将这两段组装成一个无环或者有环单链表
这个例子环的入口结点的结点值为3，所以返回结点值为3的结点

## 思路

* 快慢指针，和判断链表中是否有环是一回事儿
* 简单说就是如果有环，两个人速度不一样跑着跑着是一定会套圈的
* 相遇点不等于是环的入口，画图比划
* 【头结点到入口距离等于相遇点到入口的距离】
* 那么在相遇时，将快指针fast重新放到头结点A
* 慢指针slow的位置不变，且快指针的速度改为与慢指针slow相同
* 那么快指针fast从头结点A走过X路程后到达环的入口结点
* 慢指针slow从快慢指针相遇节点走过x路程后也到达了环的入口结点
* 也即他们再次相遇时的节点就是环的入口结点

## 实现

```go
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
```

------

# 22. 链表中倒数最后k个结点

## 题目

输入一个长度为 n 的链表，设链表中的元素的值为 ai ，返回该链表中倒数第k个节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。

要求：空间复杂度 O(n)，时间复杂度 O(n)
进阶：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

输入：{1,2,3,4,5},2

返回值：{4,5}

说明：返回倒数第2个节点4，系统会打印后面所有的节点来比较。

## 思路

> 快慢指针

* 第一个指针先移动k步
* 然后第二个指针再从头开始
* 这个时候这两个指针同时移动
* 当第一个指针到链表的末尾的时候
* 返回第二个指针即可

## 实现

```go
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	var fast *ListNode = pHead
	var slow *ListNode = pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			var empty *ListNode
			return empty
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
```

------

# 76. 删除链表中重复的结点

## 题目

在一个排序的链表中，存在重复的结点，请删除该链表中重复的结点，重复的结点不保留，返回链表头指针。
进阶：空间复杂度 O(n)，时间复杂度 O(n)

## 示例

例如，链表 1->2->3->3->4->4->5  处理后为 1->2->5

## 思路

在遍历单链表的时候，检查当前节点与下一点是否为相同值，如果相同，继续查找祥同值的最大长度，然后指针改变指向

## 实现

```go
func deleteDuplication(pHead *ListNode) *ListNode {
	var vhead *ListNode = &ListNode{
		Val:  -1,
		Next: pHead,
	}
	var pre *ListNode = vhead
	var cur *ListNode = pHead
	for {
		if cur == nil {
			break
		}
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
			for {
				if cur.Next == nil || cur.Val != cur.Next.Val {
					break
				}
				cur = cur.Next
			}
			cur = cur.Next
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}

	}
	return vhead.Next
}
```

------

# 18. 删除链表的节点

## 题目

给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。

1.此题对比原题有改动
2.题目保证链表中节点的值互不相同
3.该题只会输出返回的链表和结果做对比，所以若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点

数据范围:
0<=链表节点值<=10000
0<=链表长度<=10000

## 示例

输入：{2,5,1,9},5

返回值：{2,1,9}

说明：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 2 -> 1 -> 9 

## 思路

这还要啥思路

## 实现

```go
func deleteNode(head *ListNode, val int) *ListNode {
	var cur *ListNode = head
	var pre *ListNode = head
	for {
		if cur == nil {
			break
		}
		if cur.Val == val {
			if cur == head {
				head = head.Next
			}
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return head
}
```