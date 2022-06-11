package question63

import "fmt"

// 买卖股票的最好时机(一)
func Run() string {
	prices := []int{8, 9, 2, 5, 4, 7, 1}
	res := maxProfit(prices)
	fmt.Println(res)
	return "done!"
}

func maxProfit(prices []int) int {
	earn := 0        // 最大收益
	min := prices[0] // 股票对低值
	for k, v := range prices {
		if v < min {
			min = v
		}
		if k == 0 {
			// 第一天没有收入
			earn = 0
		} else {
			// 当天股票值减去最低股票值就是今天的预期最大收益，再和历史的最大收益值比较，选择较大的
			if (v - min) > earn {
				earn = v - min
			}
		}
	}
	return earn
}
