package places

import (
	"encoding/json"
)

type Places struct {
	places  []Place
	storage Storage
	api     Api
}

func NewPlaces(s Storage, a Api) *Places {
	return &Places{
		places:  []Place{},
		storage: s,
		api:     a,
	}
}

func (places *Places) AddPlace(place Place) {
	places.places = append(places.places, place)
}

func (places *Places) String() string {
	value, err := json.MarshalIndent(places.places, "", "\t")
	if err != nil {
		panic(err)
	}

	return string(value)
}

func (places *Places) Save() {
	places.storage.Open()
	defer places.storage.Close()

	places.storage.Append(places.String())
	places.storage.Append("\n")
}

func (places *Places) GetPlace(placeId string) Place {
	return places.api.GetPlace(placeId)
}
