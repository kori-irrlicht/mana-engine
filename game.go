package mana

import (
	"time"
)

// Defines how often Game.Update() is called by Run.
// Defaults to a target time of 16666666ns, which is means
// Update is called around 60 times per second => 60 FPS
//
// Overwrite this, if you want your game to run with an different frame rate
var TargetTimePerFrame = 1 * time.Second / 60

// The function used to get the current time in the gameloop
// Defaults to time.Now and normally does not need to be changed
var Timer = time.Now

type Game interface {
	Input()
	Update()
	Render(float32)
	Running() bool
}

// Run starts the gameloop
// It only returns when game.Running() is false
func Run(game Game) {

	lastTime := Timer()
	lag := 0 * time.Millisecond
	diff := 0 * time.Millisecond
	for game.Running() {

		current := Timer()
		diff = current.Sub(lastTime)
		lastTime = current
		lag += diff
		game.Input()
		for lag >= TargetTimePerFrame {
			game.Update()
			lag -= TargetTimePerFrame
		}

		game.Render(float32(lag) / float32(TargetTimePerFrame))
	}

}
