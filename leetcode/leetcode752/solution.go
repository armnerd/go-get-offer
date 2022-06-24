package leetcode752

// 打开转盘锁
func Run() string {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	OpenLock(deadends, "0202")
	return "done!"
}

func OpenLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum := strToInt(target)
	/*
		为什么是 10000？
		因为有 10000 种组合
	*/
	visited := make([]bool, 10000)
	visited[0] = true
	// 标记死锁位置
	for _, deadend := range deadends {
		num := strToInt(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}
	depth := 0             // 初始化步数
	curDepth := []int16{0} // 当前选择列表，0 代表了 "0000" 这个位置， 从这里开始
	var nextNum int16
	// 只要还有可能就继续
	for len(curDepth) > 0 {
		nextDepth := make([]int16, 0) // 预备选择列表
		for _, curNum := range curDepth {
			// 找到了
			if curNum == targetNum {
				return depth
			}
			// 模拟四个转轮，从左到右，依次尝试
			for incrementer := int16(1000); incrementer > 0; incrementer /= 10 {
				digit := (curNum / incrementer) % 10 // 获得这一位上的值
				// 往上拨动
				if digit == 9 {
					nextNum = curNum - 9*incrementer
				} else {
					nextNum = curNum + incrementer
				}
				// 不是死锁位或者没走过
				if !visited[nextNum] {
					visited[nextNum] = true // 标记走过了
					nextDepth = append(nextDepth, nextNum)
				}
				// 往下拨动
				if digit == 0 {
					nextNum = curNum + 9*incrementer
				} else {
					nextNum = curNum - incrementer
				}
				// 反正是能走
				if !visited[nextNum] {
					visited[nextNum] = true // 避免走回头路
					nextDepth = append(nextDepth, nextNum)
				}
			}
		}
		curDepth = nextDepth
		depth++
	}
	return -1
}

// 保持进位字符串转数字
func strToInt(str string) int16 {
	return int16(str[0]-'0')*1000 + int16(str[1]-'0')*100 + int16(str[2]-'0')*10 + int16(str[3]-'0')
}
