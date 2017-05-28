package places

import (
	"encoding/json"
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
)

type Places struct {
	Places  []place.Place
	storage Storage
}

func NewPlaces(s Storage) *Places {
	return &Places{
		Places:  []place.Place{},
		storage: s,
	}
}

func (places *Places) String() string {
	value, err := json.MarshalIndent(places, "", "\t")
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
