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
	} `json:"listener"`
	Backend struct {
		Address string `json:"address"`
		Prefix  string `json:"prefix"`
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
