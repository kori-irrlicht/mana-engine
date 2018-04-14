package asset

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/ttf"
)

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
	_ Loader = &TrueTypeFontLoader{}
)
