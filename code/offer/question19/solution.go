package question19

import "fmt"

// 正则表达式匹配
func Run() string {
	res := match("aaa", "a*a")
	fmt.Println(res)
	return "done!"
}

func match(str string, pattern string) bool {
	sSize := len(str)
	pSize := len(pattern)

	dp := make([][]bool, sSize+1)
	for i := range dp {
		dp[i] = make([]bool, pSize+1)
	}

	/* dp[i][j] 代表了 s[:i] 能否与 p[:j] 匹配 */

	dp[0][0] = true
	/**
	 * 根据题目的设定， "" 可以与 "a*b*c*" 相匹配
	 * 所以，需要把相应的 dp 设置成 true
	 */
	for j := 1; j < pSize && dp[0][j-1]; j += 2 {
		if pattern[j] == '*' {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < sSize; i++ {
		for j := 0; j < pSize; j++ {
			if pattern[j] == '.' || pattern[j] == str[i] {
				/* p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上 */
				dp[i+1][j+1] = dp[i][j]
			} else if pattern[j] == '*' {
				/* 此时，p[j] 的匹配情况与 p[j-1] 的内容相关。 */
				if pattern[j-1] != str[i] && pattern[j-1] != '.' {
					/**
					 * p[j] 无法与 s[i] 匹配上
					 * p[j-1:j+1] 只能被当做 ""
					 */
					dp[i+1][j+1] = dp[i+1][j-1]
				} else {
					/**
					 * p[j] 与 s[i] 匹配上
					 * p[j-1;j+1] 作为 "x*", 可以有三种解释
					 */
					dp[i+1][j+1] = dp[i+1][j-1] || /* "x*" 解释为 "" */
						dp[i+1][j] || /* "x*" 解释为 "x" */
						dp[i][j+1] /* "x*" 解释为 "xx..." */
				}
			}
		}
	}

	return dp[sSize][pSize]
}
