package input

type Key int

type Controller interface {
	IsDown(Key) bool
	Update()
}
