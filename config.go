package main

import (
	"log"
	"os"
	"os/user"

	"github.com/pelletier/go-toml/v2"
)

type Archive struct {
	Email string
	Owner string
}

type Config struct {
	Archive Archive
}

const configDir = ".yal/"
const configFileName = "config.toml"
const configFilePath = configDir + configFileName

func GetConfig() *Config {
	throw := func(err error) {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	c := Config{}

	dat, err := os.ReadFile(configFilePath)
	if err != nil {
		throw(err)
	}

	err = toml.Unmarshal(dat, &c)
	if err != nil {
		throw(err)
	}

	return &c
}

func InitConfig() error {
	var username string
	// get current user name to use as default owner
	user, err := user.Current()
	if err != nil {
		username = ""
	} else {
		username = user.Username
	}

	// create default config and turn it into toml
	config := Config{Archive: Archive{Owner: username}}
	b, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	// create config directory if does not exist
	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return err
	}

	// write default config to file, fail if config file already exists
	f, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
