# 基础

* 排序
* 遍历树

## 排序

> 稳定性

* 一组数据里有两个 x，经过某种排序算法排序之后
* 如果两个 x 的前后顺序没有改变，那我们就把这种排序算法叫作稳定的排序算法
* 如果前后顺序发生变化，那对应的排序算法就叫作不稳定的排序算法
* 多个字段排序时，稳定的排序可以保证之前其他字段排序的结果不被打乱

##### 冒泡

> 稳定

* 两层嵌套遍历，逐个对比置换

```go
func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
```

##### 插入

> 稳定

* 一个有序的数组，我们往里面添加一个新的数据后，如何继续保持数据有序呢？很简单，我们只要遍历数组，找到数据应该插入的位置将其插入即可
* 将数组的第一个元素当做【有序的数组】，遍历后边的数据，依次插入挪动从而保持有序

```go
func insertSort(arr []int) {
	l := len(arr)
	for i := 1; i < l; i++ {
		for v := 0; v < i; v++ {
			if arr[v] > arr[i] {
				arr[v], arr[i] = arr[i], arr[v]
			}
		}
	}
}
```

##### 选择

> 不稳定

* 选择排序算法的实现思路有点类似插入排序，也分已排序区间和未排序区间。但是选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾

```go
func SelectSort(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]
	}
}
```

##### 快排

> 不稳定

* 从排序数组中选择 start 到 end 之间的任意一个数据作为 pivot（分区点）
* 遍历 start 到 end 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间
* 特点：由上到下
* 优点：原地排序，节省内存
* 缺点：不稳定排序

```go
func quickSort(data []int, start, end int) {
	if start >= end {
		return
	}
	pivot := partition(data, start, end)
	quickSort(data, start, pivot-1)
	quickSort(data, pivot+1, end)
}

func partition(data []int, start int, end int) int {
	pivotValue := data[end] // 以最后一个作为分区点
	pivotPosition := start  // 已处理区间下一位【即分界线】，初始为开始位置
	cur := start
	// 遍历除结尾的所有元素
	for cur < end {
		// 找到小于分区点值的元素后，将其追加在【已处理区间】，分界线向后移
		if data[cur] < pivotValue {
			data[pivotPosition], data[cur] = data[cur], data[pivotPosition]
			pivotPosition += 1
		}
		cur++
	}
	// 分区点值放在分界线位置
	data[pivotPosition], data[end] = data[end], data[pivotPosition]
	// 返回分界线位置
	return pivotPosition
}
```

##### 归并

> 稳定

* 分治，顾名思义，就是分而治之，将一个大问题分解成小的子问题来解决。小的子问题解决了，大问题也就解决了
* 先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了
* 特点：由下到上
* 优点：稳定排序
* 缺点：占用太多内存

```go
func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	var temp = []int{}
	// 设置左右两段的起点下标
	var left, right = start, mid + 1
	// 相当于两个血条在对比中慢慢减少
	for left <= mid && right <= end {
		if arr[left] > arr[right] {
			temp = append(temp, arr[right])
			right++
		} else {
			temp = append(temp, arr[left])
			left++
		}
	}
	// 处理左边剩下的
	if left <= mid {
		temp = append(temp, arr[left:mid+1]...)
	}
	// 处理右边剩下的
	if right <= end {
		temp = append(temp, arr[right:end+1]...)
	}
	// 将临时数组里的归还回原始数组
	for pos, item := range temp {
		arr[start+pos] = item
	}
}
```

## 遍历树

* 叶子节点全都在最底层，除了叶子节点之外，每个节点都有左右两个子节点，这种二叉树就叫做满二叉树
* 叶子节点都在最底下两层，最后一层的叶子节点都靠左排列，并且除了最后一层，其他层的节点个数都要达到最大，这种二叉树叫做完全二叉树

##### 前序

> 先打印这个节点，然后再打印它的左子树，最后打印它的右子树

* 第一个是根

```go
func preOrder(root * Tree) {
	if root == nil {
		return
	}
	fmt.Println(root)
	preOrder(root.left)
	preOrder(root.right)
}
```

##### 中序

> 先打印它的左子树，然后再打印它本身，最后打印它的右子树

* 如果是二叉搜索树的话，中序遍历后的结果就是排好序的，从小到大

```go
func inOrder(root *Tree) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Println(root)
	inOrder(root.right)
}
```

##### 后序

> 先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身

* 最后一个是根
* 适合用于回溯

```go
func postOrder(root *Tree) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root)
}
```

------

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

# 树

* 二叉树的深度
* 按之字形顺序打印二叉树
* 二叉搜索树的第k个结点
* 重建二叉树
* 树的子结构
* 二叉树的镜像
* 从上往下打印二叉树
* 二叉树中和为某一值的路径(一)
* 判断是不是平衡二叉树
* 对称的二叉树
* 把二叉树打印成多行
* 在二叉树中找到两个节点的最近公共祖先
* 二叉搜索树的最近公共祖先

# 55. 二叉树的深度

## 题目

输入一棵二叉树，求该树的深度。从根结点到叶结点依次经过的结点（含根、叶结点）形成树的一条路径，最长路径的长度为树的深度，根节点的深度视为 1 。

进阶：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

输入：{1,2,3,4,5,#,6,#,#,7}

返回值：4

## 思路

* max( 头结点左子树的最大深度, 头结点右子树的最大深度)+1

## 实现

```go
func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		// 说明递归到根部了返回 0
		return 0
	}
	lval := TreeDepth(pRoot.Left)
	rval := TreeDepth(pRoot.Right)
	// 选择左右路径里更长的并算上自己【+1】
	return max(lval, rval) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

------

# 77. 按之字形顺序打印二叉树

## 题目

给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

要求：空间复杂度：O(n)，时间复杂度：O(n)

## 示例

输入：{1,2,3,#,#,4,5}

返回值：[[1],[3,2],[4,5]]

说明：如题面解释，第一层是根节点，从左到右打印结果，第二层从右到左，第三层从左到右。  

## 思路

* 从根节点开始，将左右子节点装入数组待下一轮处理
* 因为要之字形打印所以要记录方向，每轮反转

## 实现

```go
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	result := [][]int{{pRoot.Val}}
	reverse := true // 方向标识
	for len(nodeToBeMachined) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		row := make([]int, 0)
		for _, item := range nodeToBeMachined {
			// 待处理节点始终是从左向右的
			if item.Left != nil {
				nextLevelNodes = append(nextLevelNodes, item.Left)
				row = append(row, item.Left.Val)
			}
			if item.Right != nil {
				nextLevelNodes = append(nextLevelNodes, item.Right)
				row = append(row, item.Right.Val)
			}
		}
		nodeToBeMachined = nextLevelNodes
		length := len(row)
		if length != 0 {
			if reverse {
				for i := 0; i < length/2; i++ {
					row[length-1-i], row[i] = row[i], row[length-1-i]
				}
			}
			result = append(result, row)
		}
		reverse = !reverse // 转向
	}
	return result
}
```

------

# 54. 二叉搜索树的第k个结点

## 题目

给定一棵结点数为 n 二叉搜索树，请找出其中的第 k 小的TreeNode结点。

要求：空间复杂度 O(1)，时间复杂度 O(n)

注意：不是返回结点的值

## 示例

输入：{5,3,7,2,4,6,8},3

返回值：4

说明：按结点数值升序顺序可知第三小结点的值为4 ，故返回对应结点值为4的结点即可

## 思路

* 遍历树，所有节点放在数组里，排序取值
* 中序遍历顺序即为从小到大的访问顺序

> 为什么？

* 因为二叉搜索树的特点就是： 左节点 > 根节点 > 右节点
* 中序遍历又是先打印它的左子树，然后再打印它本身，最后打印它的右子树

## 实现

```go
var res *TreeNode

func KthNode(pRoot *TreeNode, k int) *TreeNode {
	inOrder(pRoot, &k)
    if res != nil {
        res.Left = nil
        res.Right = nil
    }
	return res
}

func inOrder(root *TreeNode, k *int) {
	if root == nil {
		return
	}
	inOrder(root.Left, k)
	*k -= 1
	if *k == 0 {
		res = root
		return
	}
	inOrder(root.Right, k)
}
```

------

# 7. 重建二叉树

## 热身

> 前序遍历 [ 根左右 ]

* 对于树中的任意节点来说，先打印这个节点，然后再打印它的左子树，最后打印它的右子树。

> 中序遍历 [ 左根右 ]

* 对于树中的任意节点来说，先打印它的左子树，然后再打印它本身，最后打印它的右子树。

> 后序遍历 [ 左右根 ]

* 对于树中的任意节点来说，先打印它的左子树，然后再打印它的右子树，最后打印这个节点本身。

## 题目

输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。例如输入前序遍历序列 {1,2,4,7,3,5,6,8} 和中序遍历序列 {4,7,2,1,5,3,8,6} ，则重建二叉树并返回。

## 思路

* 用前序遍历找到根结点
* 用根结点在中序遍历中切开左右子树，递归重建二叉树

## 实现

```go
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	// 找到根节点
	root_val := pre[0]
	root_node := &TreeNode{
		Val: root_val,
	}

	// 切开中序
	var index int
	for index = range vin {
		if vin[index] == root_val {
			break
		}
	}

	// 大树拆小树递归处理
	root_node.Left = reConstructBinaryTree(pre[1:1+index], vin[:index])
	root_node.Right = reConstructBinaryTree(pre[1+index:], vin[index+1:])

	return root_node
}
```

------

# 26. 树的子结构

## 题目

输入两棵二叉树A，B，判断B是不是A的子结构。（我们约定空树不是任意一个树的子结构）

## 示例

输入：{8,8,7,9,2,#,#,#,#,4,7},{8,9,2}

返回值：true

## 思路

* 第一步在树A中找到和B的根节点的值一样的结点R
* 第二步再判断树A中以R为根结点的子树是不是包含和树B一样的结构

## 实现

```go
func HasSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
	if pRoot1 == nil || pRoot2 == nil {
		return false
	}
	if pRoot1.Val == pRoot2.Val {
		if IsSubtree(pRoot1, pRoot2) {
			return true
		}
	}
	return HasSubtree(pRoot1.Left, pRoot2) || HasSubtree(pRoot1.Right, pRoot2)
}

func IsSubtree(pRoot1 *TreeNode, pRoot2 *TreeNode) bool {
    // 先判断 pRoot2
    if pRoot2 == nil {
		return true
	} else if pRoot1 == nil {
        return false
    }
	if pRoot1.Val != pRoot2.Val {
		return false
	}
	return IsSubtree(pRoot1.Left, pRoot2.Left) && IsSubtree(pRoot1.Right, pRoot2.Right)
}
```

------

# 27. 二叉树的镜像

## 题目

操作给定的二叉树，将其变换为源二叉树的镜像
要求： 空间复杂度 O(n)。本题也有原地操作，即空间复杂度 O(1) 的解法，时间复杂度 O(n)

## 示例

输入：{8,6,10,5,7,9,11}

返回值：{8,10,6,11,9,7,5}

## 思路

* 递归置换

## 实现

```go
func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	temp := pRoot.Left
	pRoot.Left = pRoot.Right
	pRoot.Right = temp
	Mirror(pRoot.Left)
	Mirror(pRoot.Right)
	return pRoot
}
```

------

# 32. 从上往下打印二叉树

## 题目

不分行从上往下打印出二叉树的每个节点，同层节点从左至右打印。例如输入{8,6,10,#,#,2,1}，如以下图中的示例二叉树，则依次打印8,6,10,2,1(空节点不打印，跳过)，请你将打印的结果存放到一个数组里面，返回

## 示例

输入：{8,6,10,#,#,2,1}

返回值：[8,6,10,2,1]

## 思路

* 类似于【把二叉树打印成多行】

## 实现

```go
var list = []int{}

func PrintFromTopToBottom(pRoot *TreeNode) []int {
	if pRoot == nil {
		return nil
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	for len(nodeToBeMachined) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		for _, item := range nodeToBeMachined {
			// 先把本节点的放入
			list = append(list, item.Val)
			// 把下层节点先收着下轮处理
			if item.Left != nil {
				nextLevelNodes = append(nextLevelNodes, item.Left)
			}
			if item.Right != nil {
				nextLevelNodes = append(nextLevelNodes, item.Right)
			}
		}
		nodeToBeMachined = nextLevelNodes
	}
	return list
}
```

------

# 82. 二叉树中和为某一值的路径(一)

## 题目

给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

要求：空间复杂度 O(n)，时间复杂度 O(n)
进阶：空间复杂度 O(树的高度)，时间复杂度 O(n)

## 示例

输入：{5,4,8,1,11,#,9,#,#,2,7},22

返回值：true

## 思路

> 深度优先遍历 + 回溯

* 深度优先遍历首先判断当前节点，如果为 nil 则表示递归至最深层开始回溯
* dfs 函数会将 sum 减去当前节点的值，当到了叶子节点且 sum 减到零时，说明目标路径存在

## 实现

```go
func hasPathSum(root *TreeNode, sum int) bool {
	// 遍历到了叶子节点但是该路径不符合条件
	if root == nil {
		return false
	}
	// 目标值递减
	sum -= root.Val
	// 左右节点都为空说明到了叶子节点，和也减到了0，说明找到了符合条件的路径
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	// 对左右分支进行 dfs
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```

------

# 79. 判断是不是平衡二叉树

## 题目

输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树

注：我们约定空树是平衡二叉树

## 示例

输入：{1,2,3,4,5,6,7}

返回值：true

## 思路

* 平衡二叉树是左子树的高度与右子树的高度差的绝对值小于等于1，同样左子树是平衡二叉树，右子树为平衡二叉树
* 利用后序遍历：左子树、右子树、根节点,可以先递归到叶子节点，然后在回溯的过程中来判断是否满足条件

## 实现

```go
func IsBalanced_Solution(pRoot *TreeNode) bool {
	return postOrder(pRoot) != -1
}

func postOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 如果不满足平衡二叉树的定义，则返回-1，并且如果左子树不满足条件了，直接返回-1，右子树也是如此，相当于剪枝，加速结束递归
	left := postOrder(root.Left)
	if left == -1 {
		return -1
	}
	right := postOrder(root.Right)
	if right == -1 {
		return -1
	}
	// 计算高差
	sub := 0
	max := 0
	if left > right {
		sub = left - right
		max = left
	} else {
		sub = right - left
		max = right
	}
	if sub > 1 {
		return -1
	}
	// 较长子树高度加上自己的高度
	return max + 1
}
```

------

# 28. 对称的二叉树

## 题目

给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）

要求：空间复杂度 O(n)，时间复杂度 O(n)

备注：你可以用递归和迭代两种方法解决这个问题

## 示例

输入：{1,2,2,3,4,4,3}

返回值：true

## 思路

* 以根节点为准生成两个副本递归比较
* 递归终止条件：左右节点都为空直接返回true，否则，如果只有一个为空返回false
* 对于相同位置的节点，【值相同】且【a的左枝等于b的右枝】且【a的右枝等于b的左枝】时是对称的

## 实现

```go
func isSymmetrical(pRoot *TreeNode) bool {
	return isSame(pRoot, pRoot)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val &&
		isSame(root1.Left, root2.Right) &&
		isSame(root1.Right, root2.Left)
}
```

------

# 78. 把二叉树打印成多行

## 题目

给定一个节点数为 n 二叉树，要求从上到下按层打印二叉树的 val 值，同一层结点从左至右输出，每一层输出一行，将输出的结果存放到一个二维数组中返回。

要求：空间复杂度：O(n)，时间复杂度：O(n)

## 示例

输入：{1,2,3,#,#,4,5}

返回值：[[1],[2,3],[4,5]] 

## 思路

* 从根节点开始，将左右子节点装入数组待下一轮处理

## 实现

```go
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	result := [][]int{{pRoot.Val}}
	for len(nodeToBeMachined) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		row := make([]int, 0)
		for _, item := range nodeToBeMachined {
			// 待处理节点始终是从左向右的
			if item.Left != nil {
				nextLevelNodes = append(nextLevelNodes, item.Left)
				row = append(row, item.Left.Val)
			}
			if item.Right != nil {
				nextLevelNodes = append(nextLevelNodes, item.Right)
				row = append(row, item.Right.Val)
			}
		}
		nodeToBeMachined = nextLevelNodes
		if len(row) != 0 {
			result = append(result, row)
		}
	}
	return result
}
```

------

# 86. 在二叉树中找到两个节点的最近公共祖先

## 题目

给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

要求：空间复杂度 O(1)，时间复杂度 O(n)

注：本题保证二叉树中每个节点的val值均不相同

## 示例

输入：[3,5,1,6,2,0,8,#,#,7,4],5,1

返回值：3

所以节点值为5和节点值为1的节点的最近公共祖先节点的节点值为3，所以对应的输出为3

## 思路

* 【划重点】【**普通**二叉树】【不保证大小关系】【只可以判断 nil】
* 从上往下搜索，发现目标或者找到梢节了再往上返回
* 左右节点都没有就返回nil，即代表没有找到目标
* 左右只有一侧返回，说明找到了一个目标，那么返回
* 如果左右节点都有值，那这个节点即是最近公共祖先

## 实现

```go
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	node := Core(root, p, q)
	return node.Val
}

// 递归判断 nil
func Core(root *TreeNode, p int, q int) *TreeNode {
	// 【稍节】没有找到
	if root == nil {
		return nil
	}
	// 找到了其中一个
	if p == root.Val || q == root.Val {
		return root
	}
	left := Core(root.Left, p, q)
	right := Core(root.Right, p, q)
	// 【左右】都没有找到
	if left == nil && right == nil {
		return nil
	}
	// 返回找到的一侧
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	// 左右子树上均能找到,那这个节点即是最近公共祖先
	return root
}
```

------

# 68. 二叉搜索树的最近公共祖先

## 题目

给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个结点p、q，最近公共祖先LCA(T,p,q)表示一个结点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。

## 示例

输入：{7,1,12,0,4,11,14,#,#,3,5},1,12

返回值：7

说明：节点1和节点12的最近公共祖先是7

## 思路

* 【划重点】【二叉**搜索**树】【左 > 跟 > 右】【排序或者判断 nil 都可以】
* 当节点p和节点q都在左子树时，即都小于根节点root时，需要在root的左子树查找其最近公共祖先；
* 当节点p和节点q都在右子树时，即都大于根节点root时，需要在root的右子树查找其最近公共祖先；
* 当节点p和节点q一个在根节点root的左子树，一个在其右子树，或者其中一个节点和根节点相同时，root就是其最近公共祖先

## 实现

```go
func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	node := core(root, p, q)
	return node.Val
}

// 迭代比较大小
func core(root *TreeNode, p int, q int) *TreeNode {
	if root == nil || p == root.Val || q == root.Val {
		return root
	}
	for root != nil {
		rootVal := root.Val
		if p > rootVal && q > rootVal {
			// 两个节点都在右侧
			root = root.Right
		} else if p < rootVal && q < rootVal {
			// 两个节点都在左侧
			root = root.Left
		} else {
			// 两个节点分别位于左右两侧
			return root
		}
	}
	return nil
}
```

# 队列 & 栈

* 用两个栈实现队列
* 包含min函数的栈
* 栈的压入、弹出序列
* 翻转单词序列

# 9. 用两个栈实现队列

## 题目

用两个栈来实现一个队列，完成队列的 Push 和 Pop 操作。 队列中的元素为 int 类型。

## 思路

> 如何快速晾冷一杯热水, 用两个杯子来回倒啊

* 队列是 FIFO, 栈是 FILO
* push 操作就直接往 stack1 中 push
* pop 时如果 stack2 为空，那么需要将 stack1 中的数据转移到 stack2 中，然后在对 stack2 进行pop
* 如果 stack2 不为空，直接 pop 就 ok

## 实现

```go
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
```

------

# 30. 包含min函数的栈

## 题目

定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素

进阶：栈的各个操作的时间复杂度是 O(1)，空间复杂度是 O(n)

## 示例

输入：["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]

返回值：-1,2,1,-1

## 思路

* 整一个辅助栈【可以理解为最小元素的历史快照】
* 新元素小于辅助栈栈顶push新元素
* 新元素大于等于辅助栈栈顶push辅助栈栈顶
* pop同时操作两个栈
* top操作就返回normal的栈顶
* min操作就返回minval的栈顶

## 实现

```go
var stack []int
var history []int

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
```

------

# 31. 栈的压入、弹出序列

## 题目

输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否可能为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如序列1,2,3,4,5是某栈的压入顺序，序列4,5,3,2,1是该压栈序列对应的一个弹出序列，但4,3,5,1,2就不可能是该压栈序列的弹出序列。
1.0<=pushV.length == popV.length <=1000
2.-1000<=pushV[i]<=1000
3.popV 的所有数字均在 pushV里面出现过
4.pushV 的所有数字均不相同

## 示例

输入：[1,2,3,4,5],[4,5,3,2,1]
返回值：true
说明：可以通过push(1)=>push(2)=>push(3)=>push(4)=>pop()=>push(5)=>pop()=>pop()=>pop()=>pop(), 这样的顺序得到[4,5,3,2,1]这个序列，返回true

输入：[1,2,3,4,5],[4,3,5,1,2]
返回值：false
说明：由于是[1,2,3,4,5]的压入顺序，[4,3,5,1,2]的弹出顺序，要求4，3，5必须在1，2前压入，且1，2不能弹出，但是这样压入的顺序，1又不能在2之前弹出，所以无法形成的，返回false

## 思路

* 首先注意不是一定全push完后再pop，两个操作可以交叉执行
* 直接模拟，因为弹出之前的值都会先入栈，所以用个栈来辅助

## 实现

```go
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
```

------

# 73. 翻转单词序列

## 题目

牛客最近来了一个新员工Fish，每天早晨总是会拿着一本英文杂志，写些句子在本子上。同事Cat对Fish写的内容颇感兴趣，有一天他向Fish借来翻看，但却读不懂它的意思。例如，“nowcoder. a am I”。后来才意识到，这家伙原来把句子单词的顺序翻转了，正确的句子应该是“I am a nowcoder.”。Cat对一一的翻转这些单词顺序可不在行，你能帮助他么？

进阶：空间复杂度 O(n)，时间复杂度 O(n)，保证没有只包含空格的字符串

## 示例

输入："nowcoder. a am I"

返回值："I am a nowcoder."

## 思路

从不是空格的字符开始到是空格的前一个字符结束，即为一个单词

## 实现

```go
func ReverseSentence(str string) string {
	if str == "" {
		return str
	}
	ret := ""
	tmp := ""
	isWord := false
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) != " " {
			// 合并一个单词
			tmp = string(str[i]) + tmp
			isWord = true
		} else if string(str[i]) == " " && isWord {
			// 找到一个单词，将单词合并到结果串中
			ret = ret + tmp + " "
			tmp = ""
			isWord = false
		}
	}
	if tmp != "" {
		ret += tmp
	}
	return ret
}
```

# 动态规划

* 连续子数组的最大和
* 跳台阶
* 斐波那契数列
* 矩形覆盖
* 买卖股票的最好时机(一)
* 礼物的最大价值
* 最长不含重复字符的子字符串

# 42. 连续子数组的最大和

## 题目

输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求:时间复杂度为 O(n)，空间复杂度为 O(n)
进阶:时间复杂度为 O(n)，空间复杂度为 O(1)

## 示例

输入：[1,-2,3,10,-4,7,2,-5]

返回值：18

说明：经分析可知，输入数组的子数组[3,10,-4,7,2]可以求得最大和为18 

## 思路

* 状态定义：dp[i]表示以i结尾的连续子数组的最大和。所以最终要求dp[n-1]
* 状态转移方程：dp[i] = max(array[i], dp[i-1]+array[i])

> 但是这个公式怎么来的呢？

* 假设中间两个紧挨着的位置 m 和 n， m 在 n 的前边
* 假如 m 已经和前边元素都遍历相加对比过，知道了以 m 结尾的连续子数组的最大和 为 x
* 那么以 n 结尾的连续子数组的最大和就是 n + x 和 n 中的较大值
* 当知道了以数组中每个元素结尾的子数组的连续子数组的最大和后，再取其中最大的

## 实现

```go
func FindGreatestSumOfSubArray(array []int) int {
	size := len(array)
    dp := 0
	ret := array[0]
	for i := 1; i <= size; i++ {
		dp = max(array[i-1], dp+array[i-1])
		ret = max(ret, dp)
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
```

------

# 69. 跳台阶

> 就是个斐波那契数列

## 题目

一只青蛙一次可以跳上 1 级台阶，也可以跳上 2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

## 思路

* 假设对于第 n 级台阶，总共有 f(n) 种跳法.
* 那么 f(n) = f(n-1) + f(n-2)，其中 f(1) = 1, f(2) = 2
* 其中肯定有重复运算的，搞个 map 备忘录存着

## 复杂度

* 如果单纯递归时间复杂度是 O(2^n)
* 加上记忆化搜索后时间复杂度是 O(n)，因为没有了重复计算
* 动态规划可以将空间复杂度降到 O(n) 甚至 O(1)

## 实现

```go
// 纯递归
func jumpFloor1(number int) int {
	if number <= 1 {
		return 1
	}
	return jumpFloor1(number-1) + jumpFloor1(number-2)
}

// 备忘录
func jumpFloor2(number int) int {
	note := make([]int, number)
	return fib(number, note)
}

func fib(number int, note []int) int {
	if number <= 1 {
		return 1
	}

	if note[number-1] != 0 {
		return note[number-1]
	}

	res := fib(number-1, note) + fib(number-2, note)
	note[number-1] = res
	return res
}

// 备忘录升级为动态规划, 不用递归的过程, 直接从子树求得答案, 过程是从下往上
func jumpFloor3(number int) int {
	if number <= 1 {
		return 1
	}
	note := make([]int, number)
	note[0] = 1
	note[1] = 2
	for i := 2; i < number; i++ {
		note[i] = note[i-1] + note[i-2]
	}
	return note[number-1]
}

// 究极动态规划
func jumpFloor4(number int) int {
	if number <= 1 {
		return 1
	}
	var a = 1
	var b = 1
	var c int
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
```

------

# 10. 斐波那契数列

> 这个数列从第 3 项开始，每一项都等于前两项之和

## 题目

大家都知道斐波那契数列，现在要求输入一个整数 n，请你输出斐波那契数列的第 n 项（从 0 开始，第 0 项为 0，第 1 项是 1 ）。 n <= 39

## 思路

> 递归

F(0) = 0，F(1) = 1, F(n) = F(n - 1) + F(n - 2)

## 实现

```go
func Fibonacci(n int) int {
	f1 := 0
	f2 := 1
	if n == 0 {
		return f1
	} else if n == 1 {
		return f2
	}
	for i := 1; i < n; i++ {
		tmp := f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}
```

------

# 70. 矩形覆盖

## 题目

我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，从同一个方向看总共有多少种不同的方法？

进阶：空间复杂度 O(1)，时间复杂度 O(n)

注意：约定 n == 0 时，输出 0

## 示例

比如n=3时，2*3的矩形块有3种不同的覆盖方法(从同一个方向看)：

输入描述：2*1的小矩形的总个数n

返回值描述：覆盖一个2*n的大矩形总共有多少种不同的方法(从同一个方向看)

## 思路

> 先递推下

* n=1时，显然只有一种方法
* n=2时，有2种方法
* n=3，有3中方法
* n=4,有5种方法

> 从n=3到n=4，怎么来的呢？

这里有2种情况：

* 直接在n=3的情况下，再后面中添加一个竖着的。这个很显然成立，有3种情况
* 然后横着的显然能添加到n-2的情况上，也就是在n=2后面，添加2个横着的。有2种情况

所以总结：f[n]表示2*n大矩阵的方法数
可以得出：f[n] = f[n-1] + f[n-2]，初始条件f[1] = 1, f[2] =2

## 实现

```go
func rectCover(number int) int {
	if number == 0 || number == 1 || number == 2 {
		return number
	}
	first := 1
	second := 2
	res := 0
	for i := 3; i <= number; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}
```

------

# 63. 买卖股票的最好时机(一)

## 题目

假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

要求：空间复杂度 O(1)，时间复杂度 O(n)

## 示例

输入：[8,9,2,5,4,7,1]

返回值：5

说明：在第3天(股票价格 = 2)的时候买入，在第6天(股票价格 = 7)的时候卖出，最大利润 = 7-2 = 5 ，不能选择在第2天买入，第3天卖出，这样就亏损7了；同时，你也不能在买入前卖出股票。

## 思路

* 你不能在买入股票前卖出股票
* 动态规划

## 实现

```go
func maxProfit(prices []int) int {
	earn := 0        // 最大收益
	min := prices[0] // 股票对低值
	for k, v := range prices {
		if v < min {
			min = v
		}
		if k == 0 {
			// 第一天没有收入
			earn = 0
		} else {
			// 当天股票值减去最低股票值就是今天的预期最大收益，再和历史的最大收益值比较，选择较大的
			if (v - min) > earn {
				earn = v - min
			}
		}
	}
	return earn
}
```

------

# 47. 礼物的最大价值

## 题目

在一个 m×n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
如输入这样的一个二维数组，
[
[1,3,1],
[1,5,1],
[4,2,1]
]
那么路径 1→3→5→2→1 可以拿到最多价值的礼物，价值为12

## 思路

* 每次只能向下或者向上走一格，所以对于每一个单元格来说，只能通过其左边的单元格或者上边的单元格到达，因此dp数组保存走到当前grid单元格能够得到的最大价值
* 对于第i行第j列的格子而言，走到这个单元格所能得到的最大值为 dp[同列上一行] 与 dp[同行前一列] 中最大值加上本格的值
* 递推公式： max(dp[i-1][j], dp[i][j-1]) + grid[i][j]

## 实现

```go
func maxValue(grid [][]int) int {
	// 初始化最上一行
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j-1] + grid[0][j]
	}
	// 初始化最左一列
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	// 逐列
	for i := 1; i < len(grid); i++ {
		// 逐行
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

------

# 48. 最长不含重复字符的子字符串

## 题目

请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度

## 示例

> 1

输入："abcabcbb"

返回值：3

说明：因为无重复字符的最长子串是"abc"，所以其长度为 3

> 2

输入："bbbbb"

返回值：1

说明：因为无重复字符的最长子串是"b"，所以其长度为 1

> 3

输入："pwwkew"

返回值：3

说明：因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串

## 思路

* 用dp[]记录状态，dp[i]表示以下标为i的字符结尾不包含重复字符的最长子字符串长度
* 再在dp中求maxdp
* 每次可以根据dp的前一个状态推导出后一个状态
* 因此可以省略dp数组，使用一个变量记录dp值，使用maxdp记录最大的dp值

## 实现

```go
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	prevMaxLength, maxLength := 0, 0
	for k, v := range s {
		// 从上一位开始往前数，数到以上一位结尾的最长不重复字符串的第一个字符
		i := k - 1
		for i >= (k - prevMaxLength) {
			if string(v) == string(s[i]) {
				break
			}
			i--
		}
		prevMaxLength = k - i
		// 比较历史中最大
		if prevMaxLength > maxLength {
			maxLength = prevMaxLength
		}
	}
	return maxLength
}
```

# 搜索

* 数字在升序数组中出现的次数
* 二维数组中的查找
* 旋转数组的最小数字
* 数字序列中某一位的数字

# 53. 数字在升序数组中出现的次数

## 题目

给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

要求：空间复杂度 O(1)，时间复杂度 O(logn)

## 示例

输入：[1,2,3,3,3,3,4,5],3

返回值：4

## 思路

* 暴力遍历是可以，但是面试官不开心
* 利用数组升序的特点使用二分查找
* 目标值如果有多个，肯定是连在一起
* 首先找到一个目标，再顺着两头数
* 或者可以先找上下边界
* 上界定义为：不管目标值存在与否，都指向大于目标值的第一个值
* 下界定义为：如果存在目标值，则指向第一个目标值，否则，如果不存在， 则指向大于目标值的第一个值
* 最后的结果就是：上界下标 - 上界下标

## 实现

```go
func GetNumberOfK(data []int, k int) int {
	cur := 0
	end := len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] < k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	lbound := cur
	cur = 0
	end = len(data)
	for {
		if cur >= end {
			break
		}
		mid := cur + (end-cur)/2
		if data[mid] <= k {
			cur = mid + 1
		} else {
			end = mid
		}
	}
	rbound := cur
	return rbound - lbound
}
```

------

# 4. 二维数组中的查找

## 题目

在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

```
[
 [1,2,8,9],
 [2,4,9,12],
 [4,7,10,13],
 [6,8,11,15]
]
```

给定 target = 7，返回 true。
给定 target = 3，返回 false。

## 思路

> 二分查找 / 分治

* 从右上角(或者左下角)开始查找
* 如果 target 更大，指针下移
* 如果 target 更小，指针左移

## 复杂度

* 时间复杂度：O(m+n) ，其中 m 为行数，n 为列数，最坏情况下，需要遍历 m+n 次。
* 空间复杂度：O(1)

## 实现

```go
func Find(target int, array [][]int) bool {
	// 参数检查
	var m, n int
	if m = len(array); m == 0 {
		return false
	}
	if n = len(array[0]); n == 0 {
		return false
	}

	// 行与列
	r := 0
	c := n - 1
	for {
		// 循环结束条件
		if r >= m || c < 0 {
			break
		}

		// 二分查找
		if target == array[r][c] {
			return true
		} else if target > array[r][c] {
			r++
		} else {
			c--
		}
	}
	return false
}
```

------

# 11. 旋转数组的最小数字

## 题目

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

## 示例

输入 [3,4,5,1,2]
返回值 1

## 思路

二分查找，关键在于旋转前的数组非递减排序，即递增但是会有相同的连续值

## 实现

```go
func minNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	low := 0
	high := len(rotateArray) - 1
	for {
		if low >= high {
			break
		}
		mid := (high + low) / 2
		if rotateArray[mid] > rotateArray[high] {
			// 中位大于高位，低位是中位后一位
			low = mid + 1
		} else if rotateArray[mid] < rotateArray[high] {
			// 中位小于高位，高位置为当前中位
			high = mid
		} else {
			if rotateArray[high-1] > rotateArray[high] {
				// 高位前一位大于高位，最小值即使当前高位
				low = high
				break
			}
			// 连续相同值的情况，高位向前推进一步
			high -= 1
		}
	}
	return rotateArray[low]
}
```

------

# 44. 数字序列中某一位的数字

## 题目

数字以 0123456789101112131415... 的格式作为一个字符序列，在这个序列中第 2 位（从下标 0 开始计算）是 2 ，第 10 位是 1 ，第 13 位是 1 ，以此类题，请你输出第 n 位对应的数字。

## 思路

* 首先确定该数字是属于几位数的
* 确定该数字是哪个数
* 确定是该数中哪一位

## 实现

```go
import "strconv"

func findNthDigit(n int) int {
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = start * digit * 9
	}
	num := start + (n-1)/digit
	res, _ := strconv.Atoi(string(strconv.Itoa(num)[(n-1)%digit]))
	return res
}
```

# 排序

* 数组中重复的数字
* 数组中的逆序对

# 3. 数组中重复的数字

## 题目

在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任一一个重复的数字。 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，那么对应的输出是2或者3。存在不合法的输入的话输出-1

进阶：空间复杂度 O(n)，时间复杂度 O(n)

## 示例

输入：[2,3,1,0,2,5,3]

返回值：2

说明：2或3都是对的

## 思路

使用map

## 实现

```go
func duplicate(numbers []int) int {
	res := -1
	replicateSet := make(map[int]int)
	for _, v := range numbers {
		_, ok := replicateSet[v]
		if ok {
			res = v
			break
		}
		replicateSet[v] = 1
	}
	return res
}
```

------

# 51. 数组中的逆序对

## 题目

在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

要求：空间复杂度 O(n)，时间复杂度 O(nlogn)

输入描述：题目保证输入的数组中没有的相同的数字

## 示例

输入：[1,2,3,4,5,6,7,0]

返回值：7

## 思路

* 首先【前面一个数字大于后面的数字】这俩数不一定紧挨着
* 然后，逆序是说 a[i] > a[j]，i < j
* 那么在排序的过程中，会把 a[i] 和 a[j] 交换过来，每交换一次计数 +1
* 暴力冒泡遍历不满足题目要求，需要用到归并排序
* 归并排序的过程就是，递归划分整个区间为基本相等的左右两个区间，然后就合并两个有序区间
* 在合并数组的时候，当发现右边的小于左边的时候，逆序对的个数 +1

## 实现

```go
func InversePairs(data []int) int {
	if len(data) < 2 {
		return 0
	}

	replica := make([]int, len(data))
	copy(replica, data)

	count := mergeSort(data, replica, 0, len(data)-1)
	return count % 1000000007
}

func mergeSort(data, copy []int, start, end int) int {
	if start == end {
		copy[start] = data[start]
		return 0
	}
	length := (end - start) / 2
	left := mergeSort(copy, data, start, start+length)
	right := mergeSort(copy, data, start+length+1, end)

	i := start + length
	j := end
	indexCopy := end
	count := 0

	for i >= start && j >= start+length+1 {
		if data[i] > data[j] {
			copy[indexCopy] = data[i]
			indexCopy--
			i--
			count += j - start - length
		} else {
			copy[indexCopy] = data[j]
			indexCopy--
			j--
		}
	}
	for i >= start {
		copy[indexCopy] = data[i]
		indexCopy--
        i--
	}
	for j >= start+length+1 {
		copy[indexCopy] = data[j]
		indexCopy--
        j--
	}
	return left + right + count
}
```