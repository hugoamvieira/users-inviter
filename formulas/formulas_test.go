package formulas

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSphericalLawOfCosinesEarth(t *testing.T) {
	Convey("Given valid latitude and longitude values", t, func() {
		latA := 52.654784
		lngA := -1.3426255

		latB := 45.325262
		lngB := -6.32527235

		actualDist := 891

		Convey("When func called, it should return the correct distance", func() {
			dist, err := SphericalLawOfCosinesEarth(latA, lngA, latB, lngB)
			So(err, ShouldBeNil)
			So(int(dist), ShouldEqual, actualDist)
		})
	})

	Convey("Given invalid latitude and longitude values", t, func() {
		invLatA := 102.5623135
		lngA := -1.3426255

		latB := 45.325262
		invLngB := -200.412512421

		Convey("When func called, it should return an error", func() {
			dist, err := SphericalLawOfCosinesEarth(invLatA, lngA, latB, invLngB)
			So(err, ShouldNotBeNil)
			So(dist, ShouldEqual, 0)
		})
	})
}
