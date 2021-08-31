package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	fmt.Println(uuid.New())
	fmt.Println(config)
}
