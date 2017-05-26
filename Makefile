NAME=go-chapter-1-google-places
SOURCE=cmd/main.go

all: clean depend build

run: clean build
	API_KEY=AIzaSyD7n4P7VjLkW5-mjPJVAl5YBT_JxL2gDR0 ./build/$(NAME) ./build/output.txt ChIJVXealLU_xkcRja_At0z9AGY ChIJAVkDPzdOqEcRcDteW0YgIQQ ChIJi3lwCZyTC0cRkEAWZg-vAAQ ChIJ674hC6Y_WBQRujtC6Jay33k ChIJwVPhxKtlJA0RvBSxQFbZSKY
	cat ./build/output.txt

clean:
	rm -rf build/

depend:
	go get -u -v github.com/Masterminds/glide
	glide install

build:
	go build -o build/$(NAME) $(SOURCE)

fmt:
	go fmt $(shell glide novendor)

vet:
	go vet $(shell glide novendor)

lint:
	for file in $(shell find . -name '*.go' -not -path './vendor/*'); do golint $${file}; done
