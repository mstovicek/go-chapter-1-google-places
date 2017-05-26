package main

import (
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
	"github.com/mstovicek/go-chapter-1-google-places/http/google_api"
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

	var places place.Places

	go func() {
		for place := range placesChan {
			places = append(places, place)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()

	saveToFile(os.Args[1], places)
}

func getPlace(placeId string, placesChan chan<- place.Place) {
	placesChan <- google_api.GetPlaceInformation(placeId)
}

func saveToFile(filename string, places place.Places) {
	f := file.Open(filename)
	defer f.Close()

	f.Append(places.String())
	f.Append("\n")
}
