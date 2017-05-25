package place

import "encoding/json"

type Place struct {
	PlaceId          string
	Name             string
	FormattedAddress string
	Lat              float64
	Lng              float64
}

func (place Place) ToString() string {
	value, _ := json.Marshal(place)
	return string(value)
}
