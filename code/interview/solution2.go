package interview

import (
	"fmt"
	"time"
)

/*
	题目：使用两个协程把一个字符串中的所有字符按一个数字一个字母的方式交替打印
	考察点：协程间交互，Channel
*/

func Solution2() string {
	raw := "123456abcdef"
	// raw = "123456abcdef123456"
	// raw = "123456abcdefabcdef"
	printInOrderSingle(raw)
	printInOrderMulti(raw)
	doneTest()
	return "done!"
}

func doneTest() {
	done := make(chan int)
	for i := 1; i <= 5; i++ {
		go func(sec int) {
			time.Sleep(time.Duration(sec) * time.Second)
			fmt.Println(sec)
			done <- 1
		}(i)
	}
	<-done
	<-done
	<-done
	<-done
	<-done
}

func printInOrderSingle(raw string) {
	length := len(raw)
	if length == 0 {
		return
	}
	numbers := make(chan string, length)
	chars := make(chan string, length)

	go func() {
		for k := range raw {
			v := raw[k]
			if v >= 48 && v <= 57 {
				numbers <- string(v)
			} else {
				chars <- string(v)
			}
		}
		close(numbers)
		close(chars)
	}()

	for {
		number, okN := <-numbers
		if okN {
			fmt.Println(number)
		}
		char, okC := <-chars
		if okC {
			fmt.Println(char)
		}
		if !okN && !okC {
			break
		}
	}
}

func printInOrderMulti(raw string) {
	if len(raw) == 0 {
		return
	}
	// 区分数字和字母
	numbers := make([]string, 0)
	chars := make([]string, 0)
	for k := range raw {
		v := raw[k]
		if v >= 48 && v <= 57 {
			numbers = append(numbers, string(v))
		} else {
			chars = append(chars, string(v))
		}
	}

	var length int // 选出较长作为基准
	if len(numbers) > len(chars) {
		length = len(numbers)
	} else {
		length = len(chars)
	}
	length += 1 // 这个很关键，给最后一次交接棒留出时间
	numbersTurn := make(chan int)
	charsTurn := make(chan string)
	done := make(chan int)

	go func() {
		for index := 0; index < length; index++ {
			<-numbersTurn
			if index < len(numbers) {
				fmt.Println(numbers[index])
			}
			charsTurn <- "1"
		}
		done <- 1
	}()

	go func() {
		for index := 0; index < length; index++ {
			<-charsTurn
			if index < len(chars) {
				fmt.Println(chars[index])
			}
			numbersTurn <- 1
		}
		// 因为是后执行，这里不需要 done 了
	}()

	numbersTurn <- 1 // 发令枪
	<-done           // 蚌埠住了
}
