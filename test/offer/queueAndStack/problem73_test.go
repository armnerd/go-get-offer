package queueandstack

import (
	"go-get-offer/offer/queueAndStack/question73"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem73DataSet struct {
	Str          string
	ExpectResult string
}

func TestProblem73(t *testing.T) {
	dataSet := []Problem73DataSet{
		{
			"nowcoder. a am I",
			"I am a nowcoder.",
		},
	}
	convey.Convey("翻转单词序列", t, func() {
		for _, v := range dataSet {
			res := question73.ReverseSentence(v.Str)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
