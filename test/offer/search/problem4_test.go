package search

import (
	"go-get-offer/offer/search/question4"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem4DataSet struct {
	Target       int
	Array        [][]int
	ExpectResult bool
}

func TestProblem4(t *testing.T) {
	dataSet := []Problem4DataSet{
		{
			7,
			[][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			true,
		},
		{
			1,
			[][]int{
				{2},
			},
			false,
		},
		{
			3,
			[][]int{
				{1, 2, 8, 9},
				{2, 4, 9, 12},
				{4, 7, 10, 13},
				{6, 8, 11, 15},
			},
			false,
		},
	}
	convey.Convey("二维数组中的查找", t, func() {
		for _, v := range dataSet {
			res := question4.Find(v.Target, v.Array)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
