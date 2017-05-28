package main

import (
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

	placesChan := make(chan *places.Place)

	places := places.NewPlaces(
		file.NewFile(os.Args[1]),
		google_api.NewGoogleApi(),
	)

	for _, placeId := range placeIds {
		go getPlace(places, placeId, placesChan)
	}

	go func() {
		for place := range placesChan {
			if place != nil {
				places.AddPlace(*place)
			}
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()

	places.Save()
}

func getPlace(places *places.Places, placeId string, placesChan chan<- *places.Place) {
	p := places.GetPlace(placeId)
	placesChan <- p
}
