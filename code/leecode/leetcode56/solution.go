package leetcode56

import (
	"sort"
)

// 合并区间
func Run() string {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	merge(intervals)
	return "done!"
}

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 1 {
		return intervals
	}
	// 首先根据 intervals 每个区间的 start 做升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < length-1; i++ {
		// 如果当前 end 大于下一个 start
		if intervals[i][1] >= intervals[i+1][0] {
			// 如果当前 end 小于下一个 end 使用下一个 end 值
			if intervals[i][1] < intervals[i+1][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			// 把下一个兼并
			intervals = append(intervals[:i+1], intervals[i+2:]...)
			i--      // 扭秧歌倒退一步
			length-- // 数组长度相应的减小
		}
	}
	return intervals
}
