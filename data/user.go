package data

import (
	"encoding/json"
	"errors"

	"github.com/hugoamvieira/intercom-users-inviter/formulas"
)

var (
	errNoBytesToRead   = errors.New("No bytes to read from lineBytes")
	errInvalidUserData = errors.New("User's data is invalid")
)

// User is the internal structure that defines a User in our code.
type User struct {
	ID        int64   `json:"user_id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude,string"`
	Longitude float64 `json:"longitude,string"`
}

// NewUserFromJSONBytes receives JSON bytes and tries to unmarshal it to a User structure
// It errors if you pass in an empty / nil slice, if the marshalled user's data is not valid (see valid func),
// or if the unmarshalling process fails.
func NewUserFromJSONBytes(lineBytes []byte) (*User, error) {
	if lineBytes == nil || len(lineBytes) == 0 {
		return nil, errNoBytesToRead
	}

	var user User
	err := json.Unmarshal(lineBytes, &user)
	if err != nil {
		return nil, err
	}

	if !user.valid() {
		return nil, errInvalidUserData
	}
	return &user, nil
}

// IsWithinDistanceFromCoords receives coordinates and distance in kilometers, calculates the
// great-circle distance and returns true if it is within the given distance (inclusive).
// Returns an error if given coordinates are not valid.
func (u *User) IsWithinDistanceFromCoords(lat, lng, distKm float64) (bool, error) {
	// While the haversine formula is better suited for small distances (which we will be doing),
	// the spherical law of cosines formula is a bit easier to read and it doesn't really pose
	// rouding errors for distances larger than a few meters if we use 64-bit floats (which we will).
	d, err := formulas.SphericalLawOfCosinesEarth(u.Latitude, u.Longitude, lat, lng)
	if err != nil {
		return false, err
	}

	return (d <= distKm), nil
}

func (u *User) valid() bool {
	return u.ID > 0 && // Has user id greater than 0 - I assume 0 is never possible here
		u.Name != "" && // Has a name that is not blank
		(u.Latitude >= -90.0 && u.Latitude <= 90.0) && // Has a valid latitude
		(u.Longitude >= -180.0 && u.Longitude <= 180.0) // Has a valid longitude
}
