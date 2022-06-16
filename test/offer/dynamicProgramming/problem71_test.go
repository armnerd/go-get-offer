package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question71"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem71DataSet struct {
	N            int
	ExpectResult int
}

func TestProblem71(t *testing.T) {
	dataSet := []Problem71DataSet{
		{
			3,
			4,
		},
		{
			1,
			1,
		},
	}
	convey.Convey("跳台阶扩展问题", t, func() {
		for _, v := range dataSet {
			res := question71.JumpFloorII(v.N)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
