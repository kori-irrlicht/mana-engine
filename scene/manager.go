package scene

import "fmt"

type Manager interface {
	Scene

	// Current returns the current scene
	Current() (Scene, error)

	// Register adds a new scene to the manager
	// Returns an error, if the name already exists
	Register(Name, Scene) error

	// Next requests the manager to transition to the scene with the given name
	// Returns the scene or an error, if the scene does not exist
	Next(Name) (Scene, error)

	StartWith(Name) error
}

type defaultManager struct {
	scenes  map[Name]Scene
	current Name
}

func newDefaultManager() *defaultManager {
	dm := &defaultManager{}
	dm.scenes = make(map[Name]Scene)

	return dm
}

var DefaultManager = newDefaultManager()

func (d *defaultManager) Current() (Scene, error) {
	if s, ok := d.scenes[d.current]; ok {
		return s, nil
	}
	return nil, fmt.Errorf("No current scene")
}
func (d *defaultManager) Entry()                   {}
func (d *defaultManager) Exit()                    {}
func (d *defaultManager) Input()                   {}
func (d *defaultManager) Update()                  {}
func (d *defaultManager) Render(float32)           {}
func (d *defaultManager) Next(Name) (Scene, error) { return nil, nil }
func (d *defaultManager) Register(name Name, scene Scene) error {
	if _, ok := d.scenes[name]; ok {
		return fmt.Errorf("Scene with name '%s' already registered", name)
	}
	d.scenes[name] = scene
	return nil
}
func (d *defaultManager) Ready() bool { return true }
func (d *defaultManager) StartWith(name Name) error {
	if _, ok := d.scenes[name]; !ok {
		return fmt.Errorf("No scene with name '%s'", name)
	}
	d.current = name
	return nil
}

// Force interface implementation
var _ Manager = DefaultManager
