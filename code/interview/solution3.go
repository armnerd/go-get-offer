package interview

import (
	"fmt"
)

/*
	题目：用 N 个协程（ n > =3 ）顺序打印 0 - 100
*/

func Solution3() string {
	PrintInOrder(3)
	return "done!"
}

func PrintInOrder(goroutineNum int) {
	// 创建 chan Slice
	chanSet := make([]chan int, goroutineNum)
	// 完成信号
	exitChan := make(chan bool)
	for i := 0; i < goroutineNum; i++ {
		// 创建 chan
		chanSet[i] = make(chan int)
	}
	// 给最后一个 chan 写一条数据，为了第一次输出从第 1 个 chan 输出
	go func() {
		// 导火线
		chanSet[goroutineNum-1] <- 1
	}()

	var result = 0 // 打印的初始值
	for i := 0; i < goroutineNum; i++ {
		var previousChan chan int
		var nextChan chan int
		if i == 0 {
			// 第一个协程的上一个是最后一个协程，就形成一个循环
			previousChan = chanSet[goroutineNum-1]
		} else {
			previousChan = chanSet[i-1]
		}
		nextChan = chanSet[i]
		go func(i int, previousChan, nextChan chan int) {
			for {
				// 等待上一位的信号
				<-previousChan
				if result > 100 {
					// 超过 100 就退出
					exitChan <- true
					return
				}
				fmt.Printf("goroutine%d: %d \n", i, result)
				result = result + 1
				// 给下一位信号
				nextChan <- 1
			}
		}(i, previousChan, nextChan)
	}
	// 等待完成信号
	<-exitChan
}
