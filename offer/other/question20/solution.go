package question20

import "fmt"

// 表示数值的字符串
func Run() string {
	res := isNumeric("123.45e+6")
	fmt.Println(res)
	return "done!"
}

func isNumeric(s string) bool {
	var states = []map[rune]int{
		{' ': 0, 's': 1, 'd': 2, '.': 4}, //	0
		{'d': 2, '.': 4},                 //	1
		{'d': 2, '.': 3, 'e': 6, ' ': 9}, //	2
		{'d': 5, ' ': 9, 'e': 6},         //	3
		{'d': 5},                         //	4
		{'d': 5, 'e': 6, ' ': 9},         //	5
		{'s': 7, 'd': 8},                 //	6
		{'d': 8},                         //	7
		{'d': 8, ' ': 9},                 //	8
		{' ': 9},                         //	9
	}

	state := 0
	var t rune
	for _, v := range s {
		if v == ' ' {
			t = ' '
		} else if v == '+' || v == '-' {
			t = 's'
		} else if v >= '0' && v <= '9' {
			t = 'd'
		} else if v == '.' {
			t = '.'
		} else if v == 'e' || v == 'E' {
			t = 'e'
		} else {
			t = '?'
		}
		if _, ok := states[state][t]; !ok {
			return false
		}
		state = states[state][t]
	}
	return state == 2 || state == 3 || state == 5 || state == 8 || state == 9
}
