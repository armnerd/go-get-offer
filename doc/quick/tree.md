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

给定一棵结点数为 n 二叉搜索树，请找出其中的第 k 小的 TreeNode 结点值

1. 返回第k小的节点值即可
2. 不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
3. 保证n个节点的值不一样

要求：空间复杂度 O(n)，时间复杂度 O(n)

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

func KthNode(pRoot *TreeNode, k int) int {
	inOrder(pRoot, &k)
	if res != nil {
		return res.Val
	} else {
		return -1
	}
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
func ReConstructBinaryTree(pre []int, vin []int) *TreeNode {
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
	root_node.Left = ReConstructBinaryTree(pre[1:1+index], vin[:index])
	root_node.Right = ReConstructBinaryTree(pre[1+index:], vin[index+1:])

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
