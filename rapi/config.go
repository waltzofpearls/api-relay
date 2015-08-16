package rapi

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigItem struct {
	ListenAddr    string `json: "listenAddr"`
	Downstream    string `json: "downstream"`
	ExtPathPrefix string `json: "extPathPrefix"`
	IntPathPrefix string `json: "intPathPrefix"`
}

type Config struct {
	path string
	Item *ConfigItem
}

func NewConfig() *Config {
	return &Config{}
}

func NewConfigFile(path string) *Config {
	c := NewConfig()
	c.path = path
	c.Parse()
	return c
}

func (c *Config) Parse() {
	file, _ := os.Open(c.path)
	decoder := json.NewDecoder(file)

	c.Item = new(ConfigItem)

	err := decoder.Decode(c.Item)
	if err != nil {
		panic(fmt.Sprintf("Error parsing JSON config file: %s", err))
	}
}
