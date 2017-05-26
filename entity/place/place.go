package place

import "encoding/json"

type Places []Place

type Place struct {
	PlaceId          string
	Name             string
	FormattedAddress string
	Coordinates      struct {
		Lat float64
		Lng float64
	}
}

func (places Places) String() string {
	value, err := json.MarshalIndent(places, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(value)
}
