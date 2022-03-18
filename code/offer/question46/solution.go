package question46

import (
	"fmt"
)

// 把数字翻译成字符串
func Run() string {
	res := solve("12")
	fmt.Println(res)
	return "done!"
}

func solve(nums string) int {
	if len(nums) == 0 || nums[0] == '0' {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != '0' {
			// 当前字符不等于 0 的时候, 至少有 dp[i-1] 种解法
			dp[i] = dp[i-1]
		}
		// 当前字符加上前一个字符如果是 10 到 26 之间 [包括 10 和 26] 就符合两个合并一起译码的条件
		num := (nums[i-1]-'0')*10 + (nums[i] - '0')
		if num >= 10 && num <= 26 {
			if i == 1 {
				// 若此时 i 等于 1，直接 dp[i]++
				dp[i] += 1
			} else {
				// 大于 1, 就需要加上 dp[i-2] 的结果
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(nums)-1]
}
