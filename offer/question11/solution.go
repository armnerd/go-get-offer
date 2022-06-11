package question11

import "fmt"

// 旋转数组的最小数字
func Run() string {
	rotateArray := []int{3, 6, 5, 1, 2}
	res := minNumberInRotateArray(rotateArray)
	fmt.Println(res)
	return "done!"
}

func minNumberInRotateArray(rotateArray []int) int {
	if len(rotateArray) == 0 {
		return 0
	}
	low := 0
	high := len(rotateArray) - 1
	for {
		if low >= high {
			break
		}
		mid := (high + low) / 2
		if rotateArray[mid] > rotateArray[high] {
			// 中位大于高位，低位是中位后一位
			low = mid + 1
		} else if rotateArray[mid] < rotateArray[high] {
			// 中位小于高位，高位置为当前中位
			high = mid
		} else {
			if rotateArray[high-1] > rotateArray[high] {
				// 高位前一位大于高位，最小值即使当前高位
				low = high
				break
			}
			// 连续相同值的情况，高位向前推进一步
			high -= 1
		}
	}
	return rotateArray[low]
}
