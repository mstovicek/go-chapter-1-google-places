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
		googleAPI.NewGoogleAPI(),
	)

	for _, placeID := range placeIds {
		go func(placeID string) {
			p := places.GetPlace(placeID)
			placesChan <- p
		}(placeID)
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
