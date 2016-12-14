package float

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestFloat32Zero(t *testing.T) {
	convey.Convey("float32为否为0", t, func() {
		var f float32 = 0
		convey.So(Float32Zero(f), convey.ShouldBeTrue)

		f = 0.001
		convey.So(Float32Zero(f), convey.ShouldBeFalse)
	})
}

func TestFloat32Eq(t *testing.T) {
	convey.Convey("float32应该相等", t, func() {
		var f1, f2 float32
		f1, f2 = 0.0001, 0.0001
		convey.So(Float32Eq(f1, f2), convey.ShouldBeTrue)

		f1, f2 = 0.00012, 0.0001
		convey.So(Float32Eq(f1, f2), convey.ShouldBeFalse)
	})
}
