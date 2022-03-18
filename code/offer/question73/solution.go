package question73

import "fmt"

// 翻转单词序列
func Run() string {
	raw := "nowcoder. a am I"
	res := ReverseSentence(raw)
	fmt.Println(res)
	return "done!"
}

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
