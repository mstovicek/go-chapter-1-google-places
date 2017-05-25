NAME=go-chapter-1-google-places
SOURCE=cmd/main.go

all: clean depend build

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
