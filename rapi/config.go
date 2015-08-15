package rapi

import (
	"encoding/json"
	"log"
	"os"
)

type ConfigItem struct {
	ListenAddr string `json:"listenAddr"`
	Downstream string `json:"downstream"`
}

type Config struct {
	path string
	item *ConfigItem
}

func NewConfig(path string) *Config {
	c := Config{path: path}
	c.Parse()
	return &c
}

func (c *Config) Parse() {
	file, _ := os.Open(c.path)
	decoder := json.NewDecoder(file)

	c.item = new(ConfigItem)

	err := decoder.Decode(c.item)
	if err != nil {
		log.Fatalf("Error parsing JSON config file: %s", err)
	}
}
