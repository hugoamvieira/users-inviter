package data

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewUserFromJSONBytes(t *testing.T) {
	Convey("Given valid JSON for this input", t, func() {
		id := 2
		name := "Ian McArdle"
		lat := 51.8856167
		lng := -10.4240951

		validUserJSON := fmt.Sprintf(`{"latitude": "%v", "user_id": %v, "name": "%v", "longitude": "%v"}`, lat, id, name, lng)

		Convey("It should return a user with the correct data", func() {
			user, err := NewUserFromJSONBytes([]byte(validUserJSON))
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)
			So(user.ID, ShouldEqual, id)
			So(user.Name, ShouldEqual, name)
			So(user.Latitude, ShouldEqual, lat)
			So(user.Longitude, ShouldEqual, lng)
		})
	})

	Convey("Given invalid JSON for this input", t, func() {
		id := 0
		name := "a"
		lat := 190.0
		lng := -3055.021

		invalidUserJSON := fmt.Sprintf(`{"latitude": "%v", "user_id": %v, "name": "%v", "longitude": "%v"}`, lat, id, name, lng)

		Convey("It should return an error and nil user", func() {
			user, err := NewUserFromJSONBytes([]byte(invalidUserJSON))
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, errInvalidUserData)
			So(user, ShouldBeNil)
		})
	})

	Convey("Given invalid JSON alltogether", t, func() {
		id := 0
		name := "a"
		lat := 190.0
		lng := -3055.021

		invalidUserJSON := fmt.Sprintf(`{"latitude: "%v", "user_id": %v, "name": "%v", "longitude": "%v"`, lat, id, name, lng)

		Convey("It should return an error and nil user", func() {
			user, err := NewUserFromJSONBytes([]byte(invalidUserJSON))
			So(err, ShouldNotBeNil)
			So(user, ShouldBeNil)
		})
	})

	Convey("Given an empty slice", t, func() {
		sl := make([]byte, 0)
		Convey("It should return an error and nil user", func() {
			user, err := NewUserFromJSONBytes(sl)
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, errNoBytesToRead)
			So(user, ShouldBeNil)
		})
	})

	Convey("Given nil", t, func() {
		Convey("It should return an error and nil user", func() {
			user, err := NewUserFromJSONBytes(nil)
			So(err, ShouldNotBeNil)
			So(err, ShouldEqual, errNoBytesToRead)
			So(user, ShouldBeNil)
		})
	})
}

func TestIsWithinDistanceFromCoords(t *testing.T) {
	Convey("Given valid coordinates and a user that is within the distance", t, func() {
		distKm := 100.0
		u := &User{
			ID:        1,
			Name:      "Test",
			Latitude:  53.2451022,
			Longitude: -6.238335,
		}
		lat := 53.339428
		lng := -6.257664

		Convey("Should return true and no error", func() {
			isWithin, err := u.IsWithinDistanceFromCoords(lat, lng, distKm)
			So(err, ShouldBeNil)
			So(isWithin, ShouldBeTrue)
		})
	})

	Convey("Given valid coordinates and a user that is not within the distance", t, func() {
		distKm := 100.0
		u := &User{
			ID:        1,
			Name:      "Test",
			Latitude:  23.2451022,
			Longitude: -1.238335,
		}
		lat := 53.339428
		lng := -6.257664

		Convey("Should return false and no error", func() {
			isWithin, err := u.IsWithinDistanceFromCoords(lat, lng, distKm)
			So(err, ShouldBeNil)
			So(isWithin, ShouldBeFalse)
		})
	})

	Convey("Given invalid coordinates", t, func() {
		distKm := 100.0
		u := &User{
			ID:        1,
			Name:      "Test",
			Latitude:  23.2451022,
			Longitude: -1.238335,
		}
		lat := 653.339428
		lng := -926.257664

		Convey("Should return false and error", func() {
			isWithin, err := u.IsWithinDistanceFromCoords(lat, lng, distKm)
			So(err, ShouldNotBeNil)
			So(isWithin, ShouldBeFalse)
		})
	})

	Convey("Given invalid user (this wouldn't really ever happen but tested anyways)", t, func() {
		distKm := 100.0
		u := &User{
			ID:        1,
			Name:      "Test",
			Latitude:  1523.2451022,
			Longitude: -4141411.238335,
		}
		lat := 53.339428
		lng := -6.257664

		Convey("Should return false and error", func() {
			isWithin, err := u.IsWithinDistanceFromCoords(lat, lng, distKm)
			So(err, ShouldNotBeNil)
			So(isWithin, ShouldBeFalse)
		})
	})
}
