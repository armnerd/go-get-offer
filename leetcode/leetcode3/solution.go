package leetcode3

// 无重复字符的最长子串
func Run() string {
	s := "abcabcbb"
	lengthOfLongestSubstring(s)
	return "done!"
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := len(s)
	right, res := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for right < n { // 防止越界
			if m[s[right]] != 0 {
				// 字符重复了要结束
				break
			}
			m[s[right]] = 1 // 字符出现要记录
			right++         // 不断地移动右指针
		}
		// 第 i 到 right 个字符是一个最长的无重复字符子串
		res = max(res, right-i)
	}
	return res
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
