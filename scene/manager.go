package scene

import "fmt"

type Manager interface {
	Scene

	// Current returns the current scene.
	// Returns an error if there is no current scene
	Current() (Scene, error)

	// Register adds a new scene to the manager.
	// Returns an error, if the name already exists
	Register(Name, Scene) error

	// Next requests the manager to transition to the scene with the given name.
	// Returns the scene or an error, if the scene does not exist
	Next(Name) (Scene, error)

	// StartWith sets the initial scene to the given name.
	// Returns an error if the name is not registered
	StartWith(Name) error
}

type defaultManager struct {
	Loading Scene
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
func (d *defaultManager) Entry() {}
func (d *defaultManager) Exit()  {}
func (d *defaultManager) Input() {
	s, _ := d.Current()
	s.Input()
}
func (d *defaultManager) Update() {
	s, _ := d.Current()
	s.Update()
}
func (d *defaultManager) Render(delta float32) {
	s, _ := d.Current()
	s.Render(delta)
}
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
	s, _ := d.Current()
	s.Entry()
	return nil
}

// Force interface implementation
var _ Manager = DefaultManager
