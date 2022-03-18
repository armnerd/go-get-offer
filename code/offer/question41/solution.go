package question41

import "fmt"

// 数据流中的中位数
func Run() string {
	examples := []struct {
		Nums []int
	}{
		{[]int{5, 2, 3, 4, 1, 6, 7, 0, 8}},
		{[]int{4, 5, 1, 6, 2, 7, 3, 8}},
	}
	for _, example := range examples {
		data = []int{}
		for _, v := range example.Nums {
			Insert(v)
		}
		fmt.Println(GetMedian())
	}
	return "done!"
}

var data []int

func Insert(num int) {
	if len(data) == 0 {
		data = append(data, num)
		return
	}
	for v := 0; v < len(data); v++ {
		if num < data[v] {
			num, data[v] = data[v], num
		}
	}
	data = append(data, num)
}

func GetMedian() float64 {
	length := len(data)
	if length == 0 {
		return -1
	}
	if length%2 == 0 {
		return float64(data[length/2-1]+data[length/2]) / 2
	} else {
		return float64(data[length/2])
	}
}
