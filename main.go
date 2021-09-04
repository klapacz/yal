package main

import (
	"fmt"
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

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			err := InitConfig()
			if err != nil {
				log.Fatalf("Error creating configuration file: %s", err)
			}
		case "scan":
			header, err := NewHeaderFromFile("logarion.txt")
			if err != nil {
				log.Fatalf("Error parsing header: %s", err)
			}
			log.Println(header)
		case "new":
			new()
		}
		return
	}

	config := GetConfig()
	fmt.Println(config)
}
