package question83

import (
	"fmt"
)

// 剪绳子（进阶版）
func Run() string {
	res := cutRope(4)
	fmt.Println(res)
	return "done!"
}

func CutRope(number int64) int64 {
	if number == 2 {
		return 1
	}
	if number == 3 {
		return 2
	}
	var res int64 = 1
	for number > 4 {
		res *= 3
		res %= 998244353
		number -= 3
	}
	return res * number % 998244353
}

func cutRope(number int64) int64 {
	if number == 2 {
		return 1
	}
	if number == 3 {
		return 2
	}
	cnt := number / 3
	if number%3 == 0 {
		return pow(cnt) % 998244353
	} else if number%3 == 1 {
		cnt--
		return 2 * 2 * pow(cnt) % 998244353
	} else if number%3 == 2 {
		return 2 * pow(cnt) % 998244353
	}
	return 4
}

func pow(cnt int64) int64 {
	if cnt == 0 {
		return 1
	}
	if cnt == 1 {
		return 3
	}
	part := pow(cnt / 2)
	if cnt%2 == 0 {
		return part * part % 998244353
	} else {
		return 3 * part * part % 998244353
	}
}
