package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		err := CreateConfig()
		if err != nil {
			log.Fatalf("Error creating configuration file: %s", err)
		}
		return
	}

	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	fmt.Println(uuid.New())
	fmt.Println(config)
}
