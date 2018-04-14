// Copyright 2018 kori-irrlicht
package asset

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/ttf"
)

type FileType uint

const (
	TypeTrueTypeFont = iota
)

// Holder loads assets with a given loader and stores them
type Holder interface {
	// Load loads the given file using the extra arguments(e.g. the fontsize for a fontloader)
	// and saves it internal with the given name.
	// It returns the loaded asset and any error encountered
	Load(name string, filename string, args map[string]string) (interface{}, error)

	// LoadAsync loads the given file using the extra arguments(e.g. the fontsize for a fontloader) in a goroutine
	// and saves it internal with the given name.
	// The value can be retrieved with Get(name)
	// If there an error was encountered, it can be retrieved by calling Error()
	LoadAsync(name string, filename string, args map[string]string)

	// Get retrieves a loaded asset.
	// It returns an error, if the asset is not loaded yet
	Get(name string) (interface{}, error)

	// Ready checks if all asynchronous load request are finished, but NOT if there has been an error
	Ready() bool

	// Error retrieves all encountered errors while loading asynchronous
	Error() chan error
}

// Loader contains all methods which are shared by all AssetLoader
type holder struct {
	errors  chan error
	loading chan bool
	assets  map[string]interface{}
	loader  Loader
}

// NewHolder creates an Holder from a given Loader
func NewHolder(loader Loader) Holder {
	ah := &holder{}
	ah.assets = make(map[string]interface{})
	ah.errors = make(chan error, 64)
	ah.loading = make(chan bool, 64)
	ah.loader = loader

	return ah
}

func (l *holder) Load(name, file string, args map[string]string) (interface{}, error) {
	if _, ok := l.assets[name]; ok {
		err := fmt.Errorf("Font already exists")
		return nil, err
	}

	f, err := l.loader.Load(name, file, args)
	if err != nil {
		return nil, err
	}

	l.assets[name] = f
	return f, nil
}

func (l *holder) LoadAsync(name, file string, args map[string]string) {
	l.loading <- true
	go func() {
		if _, err := l.Load(name, file, args); err != nil {
			l.errors <- err
		}
		<-l.loading
	}()
}

func (l *holder) Ready() bool {
	return len(l.loading) == 0
}

func (l *holder) Error() chan error {
	return l.Error()
}

func (l *holder) Get(name string) (interface{}, error) {
	if f, ok := l.assets[name]; !ok {
		err := fmt.Errorf("Assets with name '%s' does not exist", name)
		return nil, err
	} else {
		return f, nil
	}
}

type Loader interface {
	// Load loads the given file using the extra arguments(e.g. the fontsize for a fontloader)
	// and saves it internal with the given name.
	// It returns the loaded asset and any error encountered
	Load(name string, filename string, args map[string]string) (interface{}, error)
}

type TrueTypeFontLoader struct{}

func (l *TrueTypeFontLoader) Load(name, file string, args map[string]string) (font interface{}, err error) {
	logrus.WithFields(logrus.Fields{
		"file": file,
		"name": name,
		"args": args,
	}).Debugln("Loading font")
	size, err := strconv.Atoi(args["size"])
	if err != nil {
		return
	}
	f, err := ttf.OpenFont(file, size)
	if err != nil {
		return
	}

	return f, err
}

// Enforce interface implementation
var (
	_ Holder = &holder{}
	_ Loader = &TrueTypeFontLoader{}
)
