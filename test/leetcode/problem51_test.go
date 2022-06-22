package leetcode

import (
	"go-get-offer/leetcode/leetcode51"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem51DataSet struct {
	N            int
	ExpectResult [][]string
}

func TestProblem51(t *testing.T) {
	dataSet := []Problem51DataSet{
		{
			4,
			[][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			1,
			[][]string{
				{"Q"},
			},
		},
	}
	convey.Convey("N 皇后", t, func() {
		for _, v := range dataSet {
			res := leetcode51.SolveNQueens(v.N)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
