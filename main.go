package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const openWriteIfNotExists = os.O_WRONLY | os.O_CREATE | os.O_EXCL

const newFileName = "draft.md"

func new() {
	config := GetConfig()
	header := NewHeaderFromConfig(config)
	title := ""

	// get title from additional arguments
	if len(os.Args) > 2 {
		title = strings.Join(os.Args[2:len(os.Args)], " ")
	}

	header.Title = title

	// fail if file already exists
	f, err := os.OpenFile(newFileName, openWriteIfNotExists, 0644)
	if err != nil {
		log.Fatalf("error creating new logarion file: %s", err)
	}

	_, err = f.Write([]byte(header.ToString()))
	if err != nil {
		log.Fatalf("error writing to created logarion file: %s", err)
	}
}

func list() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("Error reading list of files: %s", err)
	}

	for _, file := range files {
		name := file.Name()
		if file.IsDir() || !strings.HasSuffix(name, ".md") {
			continue
		}

		header, err := NewHeaderFromFile(name)
		if err != nil {
			continue
		}
		fmt.Println(header.Title)
	}
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			err := InitConfig()
			if err != nil {
				log.Fatalf("Error creating configuration file: %s", err)
			}
		case "list":
			list()
		case "new":
			new()
		}
		return
	}

	config := GetConfig()
	fmt.Println(config)
}
