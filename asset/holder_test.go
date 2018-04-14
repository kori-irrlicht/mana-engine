package asset

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testLoader struct {
	inf  interface{}
	err  error
	name string
	file string
	args map[string]string
}

func (l *testLoader) Load(name, file string, args map[string]string) (inf interface{}, err error) {
	l.name = name
	l.file = file
	l.args = args
	return l.inf, l.err
}

func TestHolder(t *testing.T) {
	Convey("Given a Loader and creating a new Holder", t, func() {
		l := &testLoader{}
		h := NewHolder(l)
		Convey("The holder has the loader", func() {
			So(h.(*holder).loader, ShouldEqual, l)
		})

		Convey("Loading an asset successfully", func() {
			n := "name"
			f := "file"
			m := map[string]string{"key": "value"}
			l.inf = "inf"
			l.err = nil
			inf, err := h.Load(n, f, m)
			Convey("Holder should pass the arguments to the loader", func() {
				So(l.args, ShouldEqual, m)
				So(l.name, ShouldEqual, n)
				So(l.file, ShouldEqual, f)
			})
			Convey("Holder should return the results of the loader", func() {
				So(inf, ShouldEqual, l.inf)
				So(err, ShouldBeNil)
			})
			Convey("Loading an asset with the same name", func() {
				inf, err := h.Load(n, f, m)
				Convey("Holder should return an error", func() {
					So(inf, ShouldBeNil)
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldEqual, fmt.Errorf("Asset with name '%s' already exists", n).Error())
				})
			})
			Convey("Calling get with the asset name should return the asset", func() {
				inf2, err2 := h.Get(n)
				So(inf2, ShouldEqual, l.inf)
				So(err2, ShouldBeNil)
			})
			Convey("Holder is ready", func() {
				So(h.Ready(), ShouldBeTrue)
			})
			Convey("Holder has no error", func() {
				So(h.Error(), ShouldBeEmpty)
			})
		})
		Convey("Loading an asset not successfully", func() {
			n := "name"
			f := "file"
			m := map[string]string{"key": "value"}
			l.inf = nil
			l.err = errors.New("x")
			inf, err := h.Load(n, f, m)
			Convey("Holder should return the error", func() {
				So(inf, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err, ShouldEqual, l.err)
			})
		})

		Convey("A new Holder is ready", func() {
			So(h.Ready(), ShouldBeTrue)
		})
		Convey("A new Holder has no error", func() {
			So(h.Error(), ShouldBeEmpty)
		})

		Convey("Getting an unknown asset returns an error", func() {
			n := "unknown"
			inf, err := h.Get(n)
			So(inf, ShouldBeNil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, fmt.Errorf("Asset with name '%s' does not exist", n).Error())
		})

	})
}
