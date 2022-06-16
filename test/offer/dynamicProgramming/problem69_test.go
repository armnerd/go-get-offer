package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question69"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem69DataSet struct {
	Number       int
	ExpectResult int
}

func TestProblem69(t *testing.T) {
	dataSet := []Problem69DataSet{
		{
			2,
			2,
		},
		{
			7,
			21,
		},
	}
	convey.Convey("跳台阶", t, func() {
		for _, v := range dataSet {
			res := question69.JumpFloor(v.Number)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
