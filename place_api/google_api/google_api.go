package google_api

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
				Lat float64
				Lng float64
			}
		}
	}
}

type GoogleApi struct {
}

func NewGoogleApi() *GoogleApi {
	return &GoogleApi{}
}

func (googleApi *GoogleApi) GetPlace(placeId string) (places.Place, error) {
	var cnf config
	err := envdecode.Decode(&cnf)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	log.WithFields(log.Fields{
		"placeId": placeId,
	}).Info("Fetching place information")

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

	log.WithFields(log.Fields{
		"placeId":      placeId,
		"responseCode": resp.StatusCode,
		"status":       res.Status,
		"timeMs":       time.Since(start),
	}).Info("Fetched place information")

	if res.Status != statusOk {
		log.WithFields(log.Fields{
			"placeId": placeId,
			"status":  res.Status,
		}).Warn("Request not successful")

		return places.Place{}, fmt.Errorf("Request for place %d not successful", placeId)
	}

	return placeFromGooglePlace(res), nil
}

func placeFromGooglePlace(gPlace googlePlace) places.Place {
	return places.Place{
		PlaceId:          gPlace.Result.PlaceId,
		Name:             gPlace.Result.Name,
		FormattedAddress: gPlace.Result.FormattedAddress,
		Coordinates: places.Coordinates{
			Lat: gPlace.Result.Geometry.Location.Lat,
			Lng: gPlace.Result.Geometry.Location.Lng,
		},
	}
}
