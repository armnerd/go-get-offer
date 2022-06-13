package question16

import "fmt"

// 数值的整数次方
func Run() string {
	base := 1.23
	exponent := 2
	res := Power(base, exponent)
	fmt.Println(res)
	return "done!"
}

func Power(base float64, exp int) float64 {
	if exp == 0 {
		return 1
	}
	if exp < 0 && base == 0 {
		return -1
	}
	if exp < 0 {
		base = 1 / base
		exp = -exp
	}
	res := 1.0
	for exp != 0 {
		if exp&1 == 1 {
			res *= base
		}
		base *= base
		exp >>= 1
	}
	return res
}
