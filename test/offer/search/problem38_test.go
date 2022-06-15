package search

import (
	"go-get-offer/offer/search/question38"
	"sort"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem38DataSet struct {
	Str          string
	ExpectResult []string
}

func TestProblem38(t *testing.T) {
	dataSet := []Problem38DataSet{
		{
			"ab",
			[]string{"ab", "ba"},
		},
		{
			"aab",
			[]string{"aab", "aba", "baa"},
		},
		{
			"abc",
			[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	convey.Convey("字符串的排列", t, func() {
		for _, v := range dataSet {
			res := question38.Permutation(v.Str)
			temp := sort.StringSlice(res)
			sort.Sort(temp)
			res = []string(temp)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
