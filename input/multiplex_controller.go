package input

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

func NewMultiplexController(controller ...Controller) Controller {
	mc := &multiplexController{}
	mc.controller = controller
	return mc
}
