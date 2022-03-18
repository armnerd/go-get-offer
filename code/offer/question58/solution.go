package question58

import "fmt"

// 左旋转字符串
func Run() string {
	res := LeftRotateString("abcXYZdef", 3)
	fmt.Println(res)
	return "done!"
}

func LeftRotateString(str string, n int) string {
	if str == "" || n == len(str) {
		return str
	}
	if n > len(str) {
		n = n % len(str)
	}
	temp := []rune(str)
	res := temp[n:]
	res = append(res, temp[:n]...)
	return string(res)
}
