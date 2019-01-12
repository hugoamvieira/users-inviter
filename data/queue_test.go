package data

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEnqueueDequeue(t *testing.T) {
	Convey("Given a byte slice with data, it should add it to the queue", t, func() {
		bytes := []byte("Test")
		q := NewQueue()

		q.Enqueue(bytes)

		el, err := q.Dequeue()
		So(err, ShouldBeNil)
		So(string(el), ShouldEqual, string(bytes))
	})
}
