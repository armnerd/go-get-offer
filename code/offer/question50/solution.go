package question50

import "fmt"

// 第一个只出现一次的字符
func Run() string {
	res := FirstNotRepeatingChar("google")
	fmt.Println(res)
	return "done!"
}

func FirstNotRepeatingChar(str string) int {
	counts := make(map[rune]int)
	for _, v := range str {
		counts[v] += 1
	}
	index := -1
	for k, v := range str {
		if counts[v] == 1 {
			index = k
			break
		}
	}
	return index
}
