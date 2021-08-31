package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Archive struct {
	Email string
	Owner string
}

type Config struct {
	Archive Archive
}

const configFilePath = ".yal/config.toml"

func GetConfig() (Config, error) {
	c := Config{}

	dat, err := os.ReadFile(configFilePath)
	if err != nil {
		return c, err
	}

	err = toml.Unmarshal(dat, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
