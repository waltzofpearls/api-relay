.PHONY: test clean build

build: relay-api

test:
	go vet ./...
	go test ./...

vtest:
	go vet -v ./...
	go test -v -cover ./...

clean:
	go clean ./...

cover:
	go test -coverprofile c.out ./...
	go tool cover -html=c.out

relay-api: *.go rapi/*.go
	go build -o relay-api ./
