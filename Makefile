.PHONY: test clean build

build:relay-api

test:
	go test ./...

clean:
	go clean ./...

relay-api: *.go rapi/*.go
	go build

