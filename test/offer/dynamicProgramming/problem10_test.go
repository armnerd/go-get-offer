package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question10"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem10DataSet struct {
	N            int
	ExpectResult int
}

func TestProblem10(t *testing.T) {
	dataSet := []Problem10DataSet{
		{
			4,
			3,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
	}
	convey.Convey("斐波那契数列", t, func() {
		for _, v := range dataSet {
			res := question10.Fibonacci(v.N)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
