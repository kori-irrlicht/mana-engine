package scene

import "fmt"

// Manager contains the scenes of the game
type Manager interface {
	Managable

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
func (d *defaultManager) Input() {
	s, err := d.Current()
	if err != nil {
		panic(err.Error())
	}
	s.Input()
}
func (d *defaultManager) Update() {
	s, err := d.Current()
	if err != nil {
		panic(err.Error())
	}
	s.Update()
}
func (d *defaultManager) Render(delta float32) {
	s, err := d.Current()
	if err != nil {
		panic(err.Error())
	}
	s.Render(delta)
}

func (d *defaultManager) Next(name Name) (Scene, error) {
	s, ok := d.scenes[name]
	if !ok {
		return nil, fmt.Errorf("No scene with name '%s'", name)
	}
	curr, err := d.Current()
	if err != nil {
		panic(err.Error())
	}
	curr.Exit()

	d.current = name

	// Current() only returns an error, if there is no current scene, but
	// since we set a current scene just now, there can't be an error
	/* #nosec G104 */
	curr, _ = d.Current()
	curr.Entry()
	return s, nil
}

func (d *defaultManager) Register(name Name, scene Scene) error {
	if _, ok := d.scenes[name]; ok {
		return fmt.Errorf("Scene with name '%s' already registered", name)
	}
	d.scenes[name] = scene
	return nil
}
func (d *defaultManager) StartWith(name Name) error {
	if _, ok := d.scenes[name]; !ok {
		return fmt.Errorf("No scene with name '%s'", name)
	}
	d.current = name

	// Current() only returns an error, if there is no current scene, but
	// since we set a current scene just now, there can't be an error
	/* #nosec G104 */
	s, _ := d.Current()
	s.Entry()
	return nil
}

// Force interface implementation
var _ Manager = DefaultManager
