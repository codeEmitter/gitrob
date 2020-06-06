package common_test

import (
	. "github.com/smartystreets/goconvey/convey"
	"phantomSecrets/common"
	"testing"
)

func TestFileExists(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given a filename", t, func() {

		Convey("When the file exists", func() {
			f := "../README.md"
			b := common.FileExists(f)

			Convey("The function should return true", func() {
				So(b, ShouldEqual, true)
			})
		})

		Convey("When the file does not exist", func() {
			f := "../NOPE.md"
			b := common.FileExists(f)

			Convey("The function should return false", func() {
				So(b, ShouldEqual, false)
			})
		})
	})
}
