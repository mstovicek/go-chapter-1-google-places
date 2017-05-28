package main

import (
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
	"github.com/mstovicek/go-chapter-1-google-places/entity/places"
	"github.com/mstovicek/go-chapter-1-google-places/place_api/google_api"
	"github.com/mstovicek/go-chapter-1-google-places/storage/file"
	"os"
	"sync"
)

func main() {
	placeIds := os.Args[2:]

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(placeIds))

	placesChan := make(chan place.Place)

	for _, placeId := range placeIds {
		go getPlace(placeId, placesChan)
	}

	file := file.NewFile(os.Args[1])
	places := places.NewPlaces(file)

	go func() {
		for place := range placesChan {
			places.Places = append(places.Places, place)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()

	places.Save()
}

func getPlace(placeId string, placesChan chan<- place.Place) {
	p := google_api.GetPlaceInformation(placeId)
	placesChan <- p
}
