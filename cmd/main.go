package main

import (
	"encoding/json"
	"fmt"
	"github.com/mstovicek/go-chapter-1-google-places/http/google_api"
)

func main() {
	place := google_api.GetPlaceInformation("ChIJVXealLU_xkcRja_At0z9AGY")

	value, _ := json.Marshal(place)
	fmt.Println(string(value))
}
