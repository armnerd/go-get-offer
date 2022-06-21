package leetcode46

// 全排列
func Run() string {
	nums := []int{1, 2, 3}
	Permute(nums)
	return "done!"
}

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	track := make([]int, 0)
	onTrack := make(map[int]bool)
	var dfs func()
	dfs = func() {
		if len(track) == len(nums) {
			res = append(res, append([]int{}, track...))
			return
		}
		// nums 中不存在于 track 的元素
		for _, v := range nums {
			if onTrack[v] {
				continue
			}
			onTrack[v] = true // 搞得这么蛋疼都是因为 go 的 slice 没有 contain 之类的方法
			track = append(track, v)
			dfs()
			track = track[:len(track)-1]
			onTrack[v] = false // 还得释放
		}
	}
	dfs()
	return res
}
