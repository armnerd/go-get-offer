package leetcode322

// 零钱兑换
func Run() string {
	coins := []int{1, 2, 5}
	CoinChange(coins, 11)
	return "done!"
}

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化线性备忘录
	for k := range dp {
		dp[k] = amount + 1 // 相当于无限大，方便比较
	}
	dp[0] = 0 // 这个就是占位的
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			// 至少包含 1 枚某种硬币
			if i-coin < 0 {
				continue // 这种情况别闹
			}
			// 选择一个较小的
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1 // 木有合适的
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
