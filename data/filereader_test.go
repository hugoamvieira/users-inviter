package data

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadLinesToQueue(t *testing.T) {
	Convey("Given a valid filepath, it should put file's contents into the queue", t, func() {
		// Create file to be used in test
		filename := "tmp.tmp"
		data := []byte("Test")
		err := ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			t.Error(err)
			return
		}
		defer deleteFile(filename)

		fr := NewFileReader(filename)
		err = fr.ReadLinesToQueue()
		So(err, ShouldBeNil)

		el, err := fr.Queue.Dequeue()
		So(err, ShouldBeNil)
		So(string(el), ShouldEqual, string(data))
	})

	Convey("Given an invalid filepath, it should error and not put anything into the queue", t, func() {
		invalidFilePath := "this/should/not/exist/please/definitely/doesn't/exist/what/?"
		fr := NewFileReader(invalidFilePath)
		err := fr.ReadLinesToQueue()
		So(err, ShouldNotBeNil)

		el, err := fr.Queue.Dequeue()
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, ErrEmptyQueue)
		So(el, ShouldEqual, nil)
	})
}

func deleteFile(fp string) {
	os.Remove(fp)
}
