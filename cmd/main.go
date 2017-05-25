package main

import (
	"github.com/mstovicek/go-chapter-1-google-places/http/google_api"
	"github.com/mstovicek/go-chapter-1-google-places/storage/file"
)

func main() {
	place := google_api.GetPlaceInformation("ChIJVXealLU_xkcRja_At0z9AGY")

	file := file.Open("asd")
	defer file.Close()
	file.Append(place)
	file.Append(place)

}
