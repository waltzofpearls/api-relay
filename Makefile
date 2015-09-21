.PHONY: test clean build

build: relay-api

run: build
	./relay-api

test:
	go vet ./...
	go test ./...

clean:
	go clean ./...

relay-api: *.go rapi/*.go
	go build

