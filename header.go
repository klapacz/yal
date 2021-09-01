package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Author struct {
	Name  string
	Email string
}

type Header struct {
	ID      uuid.UUID
	Title   string
	Authors []Author
	Date    time.Time
	Topics  []string
}

const timeFormat = time.RFC3339

func (h *Header) ParseLine(text string) error {
	parts := strings.Fields(text)
	// TODO: fail if there is no second part
	switch parts[0] {
	case "ID:":
		h.ID = uuid.MustParse(parts[1])
	case "Title:":
		h.Title = strings.TrimPrefix(text, parts[0]+" ")
	case "Authors:":
		var author Author

		for i, str := range parts[1:] {
			if i%2 == 0 {
				author = Author{}
				author.Name = str
				continue
			}

			email := strings.TrimPrefix(str, "<")
			email = strings.TrimSuffix(email, ">")

			author.Email = email
			h.Authors = append(h.Authors, author)
		}
	case "Date:":
		time, err := time.Parse(timeFormat, parts[1])
		if err != nil {
			return err
		}
		h.Date = time
	case "Topics:":
		rawTopics := strings.TrimPrefix(text, parts[0]+" ")
		h.Topics = strings.Split(rawTopics, ",")
	}

	return nil
}

func GetHeader(filePath string) {
	header := Header{}

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Printf("%+v\n", header)
			return
		}
		header.ParseLine(text)
	}

	// TODO: check what code below does
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
