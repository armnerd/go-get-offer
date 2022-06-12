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