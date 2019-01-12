package formulas

import (
	"errors"
	"math"
)

var (
	errInvalidLatLng = errors.New("The latitude or longitude passed in are incorrect")
)

// SphericalLawOfCosinesEarth uses the spherical law of cosines formula to determine
// great-circle distance between two coordinate points (on earth).
// Errors if lat or lng are not valid: [-90.0, 90.0] for lat and [-180.0, 180.0] for lng.
func SphericalLawOfCosinesEarth(latA, lngA, latB, lngB float64) (float64, error) {
	if !ValidLatLng(latA, lngA) || !ValidLatLng(latB, lngB) {
		return 0.0, errInvalidLatLng
	}

	r := 6371.0 // Earth's radius, in km

	// First, convert everything to radians
	rLatA := toRadians(latA)
	rLngA := toRadians(lngA)
	rLatB := toRadians(latB)
	rLngB := toRadians(lngB)

	rLngAbsDiff := math.Abs(rLngA - rLngB)

	// Calculate central angle
	centralAngle := math.Acos(
		math.Sin(rLatA)*math.Sin(rLatB) +
			math.Cos(rLatA)*math.Cos(rLatB)*math.Cos(rLngAbsDiff),
	)

	// And then the distance (arc length) is given by multiplying
	// the central angle by earth's (in this case) radius
	distance := r * centralAngle
	return distance, nil
}

func toRadians(valDegrees float64) float64 {
	return (valDegrees * math.Pi) / 180
}

// ValidLatLng returns whether the given lat and lng coordinates are valid
func ValidLatLng(lat, lng float64) bool {
	return (lat >= -90.0 && lat <= 90.0) &&
		(lng >= -180.0 && lng <= 180.0)
}
