## Task

Write a simple api parser.

Create json array from the google places api and save it into file.

- Use google places api - https://developers.google.com/places/web-service/details
- Save the information about 5 cities to json file
- define 5 google place ids (manually)
- https://developers.google.com/maps/documentation/javascript/examples/places-placeid-finder
- delete the json file if existed before
- save the data according to the format

### JSON output format

```
[{
  long_name : "Government of Amsterdam",
  short_name : "Amsterdam",
  coordinates : {
    lat : -33.8669710,
    lng : 151.1958750
  },
  formatted_address : "Amsterdam, Netherlands",
  place_id : "D9iJyWEHuEmuEmsRm9hTkapTCrk",
  rating : 4.70,
  photo_reference: "CmRYAAAADxVKCPECPocpfBAdNZu_Sak431EBKnuEckjjoGQNhCE843NZv8HmaWmU95zCLfOqdGqq4xqLi8g_4UFHINR9xiQOUmXJhtFC_u7t3CZOX_q0MXPiIR7IJp2wHEWOZm35EhDAV0GWdK8FZzH-rKvSFZuYGhRCGnHeHEIVEDMiQerv4T1DUP6rBw"
},
…
]
```   

## Usage

Parameters of go binary:

```
go-chapter-1-google-places <output-filename> <google-place-id> <google-place-id> ...
```

Go binary requires ENV variable `API_KEY` to be set.

Install dependencies and build go binary using `Makefile`:

```
$ make all
```

Run go binary with parameters and API key:

```
API_KEY=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0 ./build/$(NAME) ./build/output.txt ChIJVXealLU_xkcRja_At0z9AGY ChIJAVkDPzdOqEcRcDteW0YgIQQ ChIJi3lwCZyTC0cRkEAWZg-vAAQ ChIJ674hC6Y_WBQRujtC6Jay33k ChIJwVPhxKtlJA0RvBSxQFbZSKY
```

or using `Makefile`

```
$ make run
```

## Output

```
$ make run
rm -rf build/
go build -o build/go-chapter-1-google-places cmd/main.go
API_KEY=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0 ./build/go-chapter-1-google-places ./build/output.txt ChIJVXealLU_xkcRja_At0z9AGY ChIJAVkDPzdOqEcRcDteW0YgIQQ ChIJi3lwCZyTC0cRkEAWZg-vAAQ ChIJ674hC6Y_WBQRujtC6Jay33k ChIJwVPhxKtlJA0RvBSxQFbZSKY
cat ./build/output.txt
[
	{
		"PlaceId": "ChIJVXealLU_xkcRja_At0z9AGY",
		"Name": "Amsterdam",
		"FormattedAddress": "Amsterdam, Netherlands",
		"Lat": 52.3702157,
		"Lng": 4.895167900000001
	},
	{
		"PlaceId": "ChIJi3lwCZyTC0cRkEAWZg-vAAQ",
		"Name": "Prague",
		"FormattedAddress": "Prague, Czechia",
		"Lat": 50.0755381,
		"Lng": 14.4378005
	},
	{
		"PlaceId": "ChIJwVPhxKtlJA0RvBSxQFbZSKY",
		"Name": "Porto",
		"FormattedAddress": "Porto, Portugal",
		"Lat": 41.1579438,
		"Lng": -8.629105299999999
	},
	{
		"PlaceId": "ChIJ674hC6Y_WBQRujtC6Jay33k",
		"Name": "Cairo",
		"FormattedAddress": "Cairo, Cairo Governorate, Egypt",
		"Lat": 30.0444196,
		"Lng": 31.2357116
	},
	{
		"PlaceId": "ChIJAVkDPzdOqEcRcDteW0YgIQQ",
		"Name": "Berlin",
		"FormattedAddress": "Berlin, Germany",
		"Lat": 52.52000659999999,
		"Lng": 13.404954
	}
]
```
