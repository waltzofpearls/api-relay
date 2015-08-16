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
	api := rapi.New(config.Item)
	api.NewEndpoint("GET", "/users")
	api.Run()
}
