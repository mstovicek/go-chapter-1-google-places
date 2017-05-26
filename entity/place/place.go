package place

import "encoding/json"

type Places []Place

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Place struct {
	PlaceId          string      `json:"place_id"`
	Name             string      `json:"name"`
	FormattedAddress string      `json:"formatted_address"`
	Coordinates      Coordinates `json:"coordinates"`
}

func (places Places) String() string {
	value, err := json.MarshalIndent(places, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(value)
}
