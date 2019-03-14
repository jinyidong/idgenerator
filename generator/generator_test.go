package generator

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenerate(t *testing.T) {
	convey.Convey("zk test",t, func() {
		convey.So(Generate(1),convey.ShouldEqual,[]int{1})
	})
}
