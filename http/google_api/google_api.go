package google_api

import (
	"encoding/json"
	"fmt"
	"github.com/joeshaw/envdecode"
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
	"io/ioutil"
	"net/http"
)

const googlePlaceEndpoint = "https://maps.googleapis.com/maps/api/place/details/json?key=%s&placeid=%s"

type config struct {
	ApiKey string `env:"API_KEY"`
}

type googlePlace struct {
	Status string `json:"status"`
	Result struct {
		PlaceId          string `json:"place_id"`
		Name             string `json:"name"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			}
		}
	}
}

func GetPlaceInformation(placeId string) place.Place {
	var cnf config
	err := envdecode.Decode(&cnf)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(fmt.Sprintf(googlePlaceEndpoint, cnf.ApiKey, placeId))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res := googlePlace{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(body)
	}

	json.Unmarshal(body, &res)

	return placeFromGooglePlace(res)
}

func placeFromGooglePlace(gPlace googlePlace) place.Place {
	return place.Place{
		PlaceId:          gPlace.Result.PlaceId,
		Name:             gPlace.Result.Name,
		FormattedAddress: gPlace.Result.FormattedAddress,
		Coordinates: struct {
			Lat float64
			Lng float64
		}{
			Lat: gPlace.Result.Geometry.Location.Lat,
			Lng: gPlace.Result.Geometry.Location.Lng,
		},
	}
}
