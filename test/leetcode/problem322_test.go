package leetcode

import (
	"go-get-offer/leetcode/leetcode322"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem322DataSet struct {
	Coins        []int
	Amount       int
	ExpectResult int
}

func TestProblem322(t *testing.T) {
	dataSet := []Problem322DataSet{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			0,
		},
	}
	convey.Convey("零钱兑换", t, func() {
		for _, v := range dataSet {
			res := leetcode322.CoinChange(v.Coins, v.Amount)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
