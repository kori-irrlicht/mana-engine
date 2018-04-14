package mana_test

import (
	"testing"

	mana "github.com/kori-irrlicht/mana-engine"
	. "github.com/smartystreets/goconvey/convey"
)

type testGame struct {
	update             bool
	render             bool
	updateBeforeRender bool
	running            chan bool
	renderCount        int
	updateCount        int
}

func (g *testGame) Update() {
	g.update = true
	g.updateCount++
	g.updateBeforeRender = !g.render
}
func (g *testGame) Render() {
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
		Convey("Running it", func() {
			g.running <- true
			g.running <- false
			mana.Run(g)
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
			Convey("Update is called twice", func() {
				So(g.updateCount, ShouldEqual, 2)
			})
			Convey("Render is called twice", func() {
				So(g.renderCount, ShouldEqual, 2)
			})
		})
	})
}
