package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "init":
			err := CreateConfig()
			if err != nil {
				log.Fatalf("Error creating configuration file: %s", err)
			}
			return
		case "scan":
			GetHeader("logarion.txt")
			return
		}
	}

	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	fmt.Println(config)
}
