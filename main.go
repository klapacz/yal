package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const newFileName = "draft.md"

func new() {
	config := GetConfig()
	header := NewHeaderFromConfig(config)
	title := ""

	if len(os.Args) > 2 {
		title = strings.Join(os.Args[2:len(os.Args)], " ")
	}

	header.Title = title

	err := os.WriteFile(newFileName, []byte(header.ToString()), 0644)
	if err != nil {
		log.Fatal(err)
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
