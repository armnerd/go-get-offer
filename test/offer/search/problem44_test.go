package search

import (
	"go-get-offer/offer/search/question44"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem44DataSet struct {
	N            int
	ExpectResult int
}

func TestProblem44(t *testing.T) {
	dataSet := []Problem44DataSet{
		{
			0,
			0,
		},
		{
			2,
			2,
		},
		{
			10,
			1,
		},
		{
			13,
			1,
		},
	}
	convey.Convey("数字序列中某一位的数字", t, func() {
		for _, v := range dataSet {
			res := question44.FindNthDigit(v.N)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
