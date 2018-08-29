package mana_test

import (
	"testing"
	"time"

	mana "github.com/kori-irrlicht/mana-engine"
	. "github.com/smartystreets/goconvey/convey"
)

type testGame struct {
	input              bool
	update             bool
	render             bool
	inputBeforeUpdate  bool
	updateBeforeRender bool
	running            chan bool
	renderCount        int
	updateCount        int
	inputCount         int
}

func (g *testGame) Input() {

	g.input = true
	g.inputCount++
	g.inputBeforeUpdate = !g.update
}

func (g *testGame) Update() {
	g.update = true
	g.updateCount++
	g.updateBeforeRender = !g.render
}
func (g *testGame) Render(float32) {
	g.render = true
	g.renderCount++
}
func (g *testGame) Running() bool {
	return <-g.running
}

func TestGame(t *testing.T) {
	Convey("Given a new Game", t, func() {
		g := &testGame{}
		g.running = make(chan bool, 5)
		mana.TargetTimePerFrame = 16 * time.Millisecond
		timeNow := time.Now()

		mana.Timer = func() time.Time {
			timeNow = timeNow.Add(16 * time.Millisecond)
			return timeNow
		}
		Convey("Running it", func() {
			g.running <- true
			g.running <- false
			mana.Run(g)
			Convey("Input is called", func() {
				So(g.input, ShouldBeTrue)
			})
			Convey("Update is called", func() {
				So(g.update, ShouldBeTrue)
			})
			Convey("Render is called", func() {
				So(g.render, ShouldBeTrue)
			})
			Convey("Update is called before Render", func() {
				So(g.updateBeforeRender, ShouldBeTrue)
			})
		})
		Convey("If game is not running", func() {
			g.running <- false
			mana.Run(g)
			Convey("Input is not called", func() {
				So(g.input, ShouldBeFalse)
			})
			Convey("Update is not called", func() {
				So(g.update, ShouldBeFalse)
			})
			Convey("Render is not called", func() {
				So(g.render, ShouldBeFalse)
			})
		})
		Convey("While game is running (2 times)", func() {
			g.running <- true
			g.running <- true
			g.running <- false
			mana.Run(g)
			Convey("Input is called twice", func() {
				So(g.inputCount, ShouldEqual, 2)
			})
			Convey("Update is called twice", func() {
				So(g.updateCount, ShouldEqual, 2)
			})
			Convey("Render is called twice", func() {
				So(g.renderCount, ShouldEqual, 2)
			})
		})

		Convey("Setting the update time to 16ms and running for 40ms", func() {
			g.running <- true
			g.running <- false
			timeNow := time.Now()

			mana.Timer = func() time.Time {
				timeNow = timeNow.Add(40 * time.Millisecond)
				return timeNow
			}
			mana.Run(g)
			Convey("Input is called once", func() {
				So(g.inputCount, ShouldEqual, 1)
			})

			Convey("Update is called twice", func() {
				So(g.updateCount, ShouldEqual, 2)

			})
			Convey("Render is called once", func() {
				So(g.renderCount, ShouldEqual, 1)
			})

		})
		Convey("Setting the update time to 16ms and running for 10ms", func() {
			g.running <- true
			g.running <- false
			timeNow := time.Now()

			mana.Timer = func() time.Time {
				timeNow = timeNow.Add(10 * time.Millisecond)
				return timeNow
			}
			mana.Run(g)

			Convey("Input is called once", func() {
				So(g.inputCount, ShouldEqual, 1)
			})
			Convey("Update is never called", func() {
				So(g.updateCount, ShouldEqual, 0)

			})
			Convey("Render is called once", func() {
				So(g.renderCount, ShouldEqual, 1)
			})

		})
	})
}
