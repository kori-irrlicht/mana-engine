package input

// Key is a virtual abstraction of physical keys on a keyboard or controller
// E.g. pressing up on a controller or keyboard could both map the same Key
type Key int

// Controller is an abstraction of a physical controller or keyboard
type Controller interface {
	// IsDown checks if the specified key is currently pressed
	IsDown(Key) bool

	// Update updates the state of the Controller.
	// If the Controller needs to process events to update its internal state,
	// Update should be called at least once per game loop cycle.
	Update()
}
