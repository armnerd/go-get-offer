package leetcode

import (
	"go-get-offer/leetcode/leetcode752"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem752DataSet struct {
	Deadends     []string
	Target       string
	ExpectResult int
}

func TestProblem752(t *testing.T) {
	dataSet := []Problem752DataSet{
		{
			[]string{"0201", "0101", "0102", "1212", "2002"},
			"0202",
			6,
		},
		{
			[]string{"8888"},
			"0009",
			1,
		},
		{
			[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
			"8888",
			-1,
		},
	}
	convey.Convey("打开转盘锁", t, func() {
		for _, v := range dataSet {
			res := leetcode752.OpenLock(v.Deadends, v.Target)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
