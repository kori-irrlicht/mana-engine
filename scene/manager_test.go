package scene

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testScene struct{}

func (s *testScene) Entry()         {}
func (s *testScene) Exit()          {}
func (s *testScene) Input()         {}
func (s *testScene) Update()        {}
func (s *testScene) Render(float32) {}
func (s *testScene) Ready() bool    { return false }

func TestDefaultManager(t *testing.T) {
	Convey("A new manager", t, func() {
		m := newDefaultManager()
		Convey("Register a scene", func() {
			var name Name = "s1"
			ts := &testScene{}
			err := m.Register(name, ts)
			Convey("There was no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Registering it again", func() {
				err := m.Register(name, ts)
				Convey("There was an error", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, fmt.Errorf("Scene with name '%s' already registered", name).Error())
				})
			})

			Convey("Selecting it as start scene", func() {
				err := m.StartWith(name)
				Convey("There was no error", func() {
					So(err, ShouldBeNil)
				})
				Convey("scene is current scene", func() {
					s, err := m.Current()
					So(s, ShouldEqual, ts)
					So(err, ShouldBeNil)
				})

			})
		})
		Convey("no current scene", func() {
			s, err := m.Current()
			So(s, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, fmt.Errorf("No current scene").Error())
		})
		Convey("Selecting an unknown start scene", func() {
			err := m.StartWith("name")
			Convey("StartWith should return an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Errorf("No scene with name 'name'").Error())

			})
		})
	})
}
