package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question19"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem19DataSet struct {
	Str          string
	Pattern      string
	ExpectResult bool
}

func TestProblem19(t *testing.T) {
	dataSet := []Problem19DataSet{
		{
			"aaa",
			"a*a",
			true,
		},
		{
			"aad",
			"c*a*d",
			true,
		},
		{
			"a",
			".*",
			true,
		},
		{
			"aaab",
			"a*a*a*c",
			false,
		},
	}
	convey.Convey("正则表达式匹配", t, func() {
		for _, v := range dataSet {
			res := question19.Match(v.Str, v.Pattern)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
