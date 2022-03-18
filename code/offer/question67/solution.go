package question67

import (
	"fmt"
	"math"
)

// 把字符串转换成整数(atoi)
func Run() string {
	res := StrToInt("82")
	fmt.Println(res)
	return "done!"
}

func StrToInt(str string) int {
	res := 0
	flag := 1
	start := 0
	// 处理空格
	for start <= len(str)-1 && str[start] == ' ' {
		start++
	}
	// 正负数
	if start <= len(str)-1 {
		if str[start] == '-' {
			flag = -1
			start++
		} else if str[start] == '+' {
			start++
		}
	}
	// 数字
	for start <= len(str)-1 && isDigital(str[start]) {
		if res == 0 {
			res += int(str[start] - '0')
			start++
		} else {
			res = res*10 + int((str[start] - '0'))
			start++
		}
		// 溢出
		if res*flag > math.MaxInt32 {
			return math.MaxInt32
		} else if res*flag < math.MinInt32 {
			return math.MinInt32
		}
	}

	res *= flag
	return res
}

func isDigital(b byte) bool {
	return b >= '0' && b <= '9'
}
