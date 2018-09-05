package scene

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testScene struct {
	input       bool
	update      bool
	render      bool
	renderDelta float32
}

func (s *testScene) Entry() {}
func (s *testScene) Exit()  {}
func (s *testScene) Input() {
	s.input = true
}
func (s *testScene) Update() {
	s.update = true
}
func (s *testScene) Render(f float32) {
	s.render = true
	s.renderDelta = f
}
func (s *testScene) Ready() bool { return false }

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
					Convey("Input is passed to the current scene", func() {
						s := s.(*testScene)
						m.Input()
						So(s.input, ShouldBeTrue)
					})
					Convey("Update is passed to the current scene", func() {
						s := s.(*testScene)
						m.Update()
						So(s.update, ShouldBeTrue)
					})
					Convey("Render is passed to the current scene", func() {
						s := s.(*testScene)
						m.Render(14)
						So(s.render, ShouldBeTrue)
						So(s.renderDelta, ShouldEqual, 14)
					})
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
		Convey("Registering multiple scenes", func() {
			var n1 Name = "s1"
			var n2 Name = "s2"
			ts := &testScene{}
			ts2 := &testScene{}
			m.Register(n1, ts)
			m.Register(n2, ts2)
			m.StartWith(n1)
			Convey("Switching to the second scene", func() {
				s, err := m.Next(n2)
				So(s, ShouldEqual, ts2)
				So(err, ShouldBeNil)
				Convey("The current scene is the 'next scene'", func() {
					curr, _ := m.Current()
					So(curr, ShouldEqual, s)
				})
			})
			Convey("Switching to a non existant scene", func() {
				s, err := m.Next("404")
				So(s, ShouldEqual, nil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "No scene with name '404'")
			})

		})
	})
}
