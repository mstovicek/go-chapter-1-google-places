package googleAPI

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/joeshaw/envdecode"
	"github.com/mstovicek/go-chapter-1-google-places/entity/places"
	"io/ioutil"
	"net/http"
	"time"
)

const googlePlaceEndpoint = "https://maps.googleapis.com/maps/api/place/details/json?key=%s&placeid=%s"

const statusOk = "OK"

type config struct {
	APIKey string `env:"API_KEY"`
}

type googlePlace struct {
	Status string `json:"status"`
	Result struct {
		PlaceID          string `json:"place_id"`
		Name             string `json:"name"`
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

type GoogleAPI struct {
	cnf config
}

func NewGoogleAPI() *GoogleAPI {
	var cnf config
	err := envdecode.Decode(&cnf)
	if err != nil {
		panic(err)
	}

	return &GoogleAPI{
		cnf: cnf,
	}
}

func (googleAPI *GoogleAPI) GetPlace(placeID string) (*places.Place, error) {
	start := time.Now()

	log.WithFields(log.Fields{
		"placeID": placeID,
	}).Info("Fetching place information")

	resp, err := http.Get(fmt.Sprintf(googlePlaceEndpoint, googleAPI.cnf.APIKey, placeID))
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

	log.WithFields(log.Fields{
		"placeID":      placeID,
		"responseCode": resp.StatusCode,
		"status":       res.Status,
		"timeMs":       time.Since(start),
	}).Info("Fetched place information")

	if res.Status != statusOk {
		log.WithFields(log.Fields{
			"placeID": placeID,
			"status":  res.Status,
		}).Warn("Request not successful")

		return nil, fmt.Errorf("Cannot get place information (place id = %s)", placeID)
	}

	p := placeFromGooglePlace(res)
	return &p, nil
}

func placeFromGooglePlace(gPlace googlePlace) places.Place {
	return places.Place{
		PlaceID:          gPlace.Result.PlaceID,
		Name:             gPlace.Result.Name,
		FormattedAddress: gPlace.Result.FormattedAddress,
		Coordinates: places.Coordinates{
			Lat: gPlace.Result.Geometry.Location.Lat,
			Lng: gPlace.Result.Geometry.Location.Lng,
		},
	}
}
