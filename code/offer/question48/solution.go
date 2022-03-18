package question48

import "fmt"

// 最长不含重复字符的子字符串
func Run() string {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
	return "done!"
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	prevMaxLength, maxLength := 0, 0
	for k, v := range s {
		// 从上一位开始往前数，数到以上一位结尾的最长不重复字符串的第一个字符
		i := k - 1
		for i >= (k - prevMaxLength) {
			if string(v) == string(s[i]) {
				break
			}
			i--
		}
		prevMaxLength = k - i
		// 比较历史中最大
		if prevMaxLength > maxLength {
			maxLength = prevMaxLength
		}
	}
	return maxLength
}
