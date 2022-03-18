package question38

// 字符串的排列
func Run() string {
	Permutation("abcde")
	return "done!"
}

func Permutation(str string) []string {
	var result []string
	if str == "" {
		return result
	}
	length := len(str)
	strByte := []byte(str)
	var handler func(int)
	handler = func(i int) {
		if i != length {
			// 用 map 排重
			strMap := make(map[string]int)
			for j := i; j < length; j++ {
				_, ok := strMap[string(strByte[j])]
				if !ok {
					strMap[string(strByte[j])] = 1
					// 把首位和剩下的依次做替换
					strByte[i], strByte[j] = strByte[j], strByte[i]
					// 把剩下的作为一个[子问题]递归处理
					handler(i + 1) // +1 决定了多锁定一位
					// 切换过的要换回来迎接下一位
					strByte[i], strByte[j] = strByte[j], strByte[i]
				}
			}
		} else {
			// 每一位都替换了才可以加入结果集
			result = append(result, string(strByte))
		}
	}
	handler(0)
	return result
}
