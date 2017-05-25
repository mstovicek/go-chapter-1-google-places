package main

import (
	"fmt"
	"github.com/mstovicek/go-chapter-1-google-places/entity/place"
	"github.com/mstovicek/go-chapter-1-google-places/http/google_api"
	"github.com/mstovicek/go-chapter-1-google-places/storage/file"
	"math/rand"
	"os"
	"time"
)

func main() {
	f := file.Open(os.Args[1])
	defer f.Close()

	for _, placeId := range os.Args[2:] {
		writePlace(f, getPlace(placeId))
	}
}

func getPlace(placeId string) place.Place {
	sleep()
	return google_api.GetPlaceInformation(placeId)
}

func writePlace(f *file.File, place place.Place) {
	f.Append(place)
}

func sleep() {
	rand.Seed(time.Now().Unix())
	ms := rand.Intn(1000)

	fmt.Printf("sleep %dms \n", ms)

	time.Sleep(time.Duration(ms) * time.Millisecond)
}
