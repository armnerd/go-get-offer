package question5

import "strings"

// 替换空格
func Run() string {
	var s = "We Are Happy"
	replaceSpace(s)
	return "done!"
}

func replaceSpace(s string) string {
	var tmp = make([]string, 0, len(s))
	for index := range s {
		if string(s[index]) == " " {
			tmp = append(tmp, "%20")
		} else {
			tmp = append(tmp, string(s[index]))
		}
	}
	return strings.Join(tmp, "")
}
