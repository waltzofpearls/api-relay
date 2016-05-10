.PHONY: test clean build

build: api-relay

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

api-relay: *.go rapi/*.go
	go build -o api-relay ./
