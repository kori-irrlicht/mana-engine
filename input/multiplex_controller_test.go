package input

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testController struct {
	down    Key
	updated bool
}

func (tc testController) IsDown(i Key) bool {
	return tc.down == i
}

func (tc *testController) Update() {
	tc.updated = true
}

// Enforce interface implementation
var _ Controller = &multiplexController{}

const (
	testKey1 Key = 1
)

func TestLogger(t *testing.T) {
	Convey("Requesting a new MultiplexController", t, func() {
		mc := NewMultiplexController()
		Convey("It is not nil", func() {
			So(mc, ShouldNotBeNil)
		})
	})
	Convey("Requesting a new MultiplexController with other controller", t, func() {
		t1 := &testController{}
		t2 := &testController{}
		mc := NewMultiplexController(t1, t2).(*multiplexController)
		Convey("It should have those controller", func() {
			So(mc.controller, ShouldResemble, []Controller{t1, t2})
		})

		Convey("T1 has key 1 down", func() {
			t1.down = testKey1
			Convey("MultiplexController has key 1 down", func() {
				So(mc.IsDown(testKey1), ShouldBeTrue)
			})
		})
		Convey("T1 has key 1 up", func() {
			Convey("MultiplexController has key 1 up", func() {
				So(mc.IsDown(testKey1), ShouldBeFalse)
			})
		})
		Convey("T2 has key 1 down", func() {
			t2.down = testKey1
			Convey("MultiplexController has a key 1 down", func() {
				So(mc.IsDown(testKey1), ShouldBeTrue)
			})
		})
		Convey("T2 has key 1 up", func() {
			Convey("MultiplexController has a key 1 up", func() {
				So(mc.IsDown(testKey1), ShouldBeFalse)
			})
		})
		Convey("T1 and T2 have key 1 up", func() {
			Convey("MultiplexController has key 1 up", func() {
				So(mc.IsDown(testKey1), ShouldBeFalse)
			})
		})

		Convey("Update MultiplexController updates subcontroller", func() {
			mc.Update()
			So(t1.updated, ShouldBeTrue)
			So(t2.updated, ShouldBeTrue)
		})
	})
}
