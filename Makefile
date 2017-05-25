NAME=go-chapter-1-google-places
SOURCE=cmd/main.go

all: clean depend build

run: clean build
	./build/$(NAME) ./build/a.txt ChIJVXealLU_xkcRja_At0z9AGY ChIJAVkDPzdOqEcRcDteW0YgIQQ ChIJi3lwCZyTC0cRkEAWZg-vAAQ
	cat ./build/a.txt

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
