package main

import "github.com/waltzofpearls/relay-api/rapi"

func main() {
	api := rapi.New("/v1").
		setListenAddr("http://localhost:8080").
		setDownstreamAddr("http://localhost:8094")
	api.NewEndpoint("GET", "/users")
	api.Run()
}
