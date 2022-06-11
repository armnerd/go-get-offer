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