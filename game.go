package mana

type Game interface {
	Update()
	Render()
	Running() bool
}

func Run(game Game) {

	for game.Running() {

		game.Update()
		game.Render()
	}

}
