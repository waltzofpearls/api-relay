package rapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Listener struct {
		Address string `json:"address"`
		Prefix  string `json:"prefix"`
		Tls     struct {
			Enable   bool   `json:"enable"`
			CertFile string `json:"certFile"`
			KeyFile  string `json:"keyFile"`
		} `json:"tls"`
	} `json:"listener"`
	Backend struct {
		Address string `json:"address"`
		Prefix  string `json:"prefix"`
		Tls     struct {
			Enable             bool `json:"enable"`
			InsecureSkipVerify bool `json:"insecureSkipVerify"`
		} `json:"tls"`
	} `json:"backend"`
}

func NewConfig() *Config {
	return &Config{}
}

func NewConfigFile(path string) *Config {
	c := NewConfig()
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(c)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON config file: %s", err))
	}

	return c
}
