package main

import (
	"flag"
	"fmt"

	"github.com/waltzofpearls/relay-api/rapi"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "c", "config.json", "Path to config file")
	flag.Parse()
	flag.Visit(func(v *flag.Flag) {
		fmt.Printf("%s - %s: %s\n", v.Usage, v.Name, v.Value)
	})

	config := rapi.NewConfigFile(configPath)
	api := rapi.New(config)

	api.NewEndpoint("POST", "/auth/login")
	api.NewEndpoint("GET", "/users")
	api.NewEndpoint("GET", "/invoices").
		TransformResponse(&InvoiceInternal{}, &InvoiceExternal{})
	api.NewEndpoint("GET", "/invoices/{Id:[A-Z0-9]+}").
		InternalPath("/invoices/{{.Id}}").
		TransformResponse(&InvoiceInternal{}, &InvoiceExternal{})

	api.Run()
}
