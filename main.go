package main

import "github.com/waltzofpearls/relay-api/rapi"

func main() {
	api := rapi.New("/v1")
	api.NewEndpoint("GET", "/users")
	api.Run()
}
