package main

import "github.com/waltzofpearls/relay-api/rapi"

func main() {
	api := rapi.New("/v1").
		setListenProtocol("http").
		setListenAddr("localhost:8080").
		setDownstreamProtocol("http").
		setDownstreamAddr("localhost:8094")
	api.NewEndpoint("GET", "/users")
	api.Run()
}
