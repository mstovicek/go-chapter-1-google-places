package place

import "encoding/json"

type Places []*Place

type Place struct {
	PlaceId          string
	Name             string
	FormattedAddress string
	Lat              float64
	Lng              float64
}

func (places Places) String() string {
	value, err := json.Marshal(places)
	if err != nil {
		panic(err)
	}

	return string(value)
}
