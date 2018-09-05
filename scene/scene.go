package scene

// Name is an identifier for a Scene
type Name string

// Scene is a part of the game, like a main menu, option, loading screen or the ingame
type Scene interface {
	Managable

	// Entry is called, if the Manager wants to switch to this scene
	// The scene is expected to load everything it needs
	Entry()

	// Ready returns true, if the scene has loaded everything it needs
	Ready() bool

	// Exit should clean up everything and is called by the manager, when a transition
	// to the next scene is requested
	Exit()
}

type Managable interface {
	// Input handles the player input
	Input()

	// Update
	Update()
	Render(float32)
}
