package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server  string
	Databse string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
