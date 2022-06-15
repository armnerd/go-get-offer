package search

import (
	"go-get-offer/offer/search/question53"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem53DataSet struct {
	Data         []int
	K            int
	ExpectResult int
}

func TestProblem53(t *testing.T) {
	dataSet := []Problem53DataSet{
		{
			[]int{1, 2, 3, 3, 3, 3, 4, 5},
			3,
			4,
		},
		{
			[]int{1, 3, 4, 5},
			6,
			0,
		},
	}
	convey.Convey("数字在升序数组中出现的次数", t, func() {
		for _, v := range dataSet {
			res := question53.GetNumberOfK(v.Data, v.K)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
