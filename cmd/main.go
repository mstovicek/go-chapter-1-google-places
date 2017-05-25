package main

import (
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
	"github.com/mstovicek/go-chapter-1-google-places/http/google_api"
	"github.com/mstovicek/go-chapter-1-google-places/storage/file"
	"os"
	"sync"
)

func main() {
	f := file.Open(os.Args[1])
	defer f.Close()

	placeIds := os.Args[2:]

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(placeIds))

	placesChan := make(chan place.Place)

	go writePlace(f, placesChan, &waitGroup)

	for _, placeId := range placeIds {
		go getPlace(placeId, placesChan)
	}

	waitGroup.Wait()
}

func getPlace(placeId string, placesChan chan<- place.Place) {
	placesChan <- google_api.GetPlaceInformation(placeId)
}

func writePlace(f *file.File, placesChan <-chan place.Place, waitGroup *sync.WaitGroup) {
	for place := range placesChan {
		f.Append(place)
		waitGroup.Done()
	}
}
