package leetcode

import (
	"go-get-offer/leetcode/leetcode1"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem1DataSet struct {
	Nums         []int
	Target       int
	ExpectResult []int
}

func TestProblem1(t *testing.T) {
	dataSet := []Problem1DataSet{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}
	convey.Convey("两数之和", t, func() {
		for _, v := range dataSet {
			res := leetcode1.TwoSum(v.Nums, v.Target)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
