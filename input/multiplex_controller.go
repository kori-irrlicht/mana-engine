package input

// multiplexController uses multiple controller to determine if a Key is pressed.
// If at least one of its internal controllers returns true for IsDown(a), the multiplexController
// also returns true for the 'a' Key
type multiplexController struct {
	controller []Controller
}

func (mc *multiplexController) IsDown(key Key) bool {
	for _, c := range mc.controller {
		if c.IsDown(key) {
			return true
		}
	}
	return false
}

func (mc *multiplexController) Update() {
	for _, c := range mc.controller {
		c.Update()
	}
}

// NewMultiplexController creates a new multiplexController from the given controller.
func NewMultiplexController(controller ...Controller) Controller {
	mc := &multiplexController{}
	mc.controller = controller
	return mc
}
