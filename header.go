package main

import (
	"bufio"
	"fmt"
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
	if len(parts) < 2 {
		return fmt.Errorf("Header field '%s' has no value", parts[0])
	}

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

func (h *Header) ToString() string {
	authors := ""
	for _, a := range h.Authors {
		authors += " " + a.Name + " " + a.Email
	}

	return "ID: " + h.ID.String() +
		"\nTitle: " + h.Title +
		"\nAuthors:" + authors +
		"\nDate: " + h.Date.Format(timeFormat) +
		"\nTopics: " + strings.Join(h.Topics, ",")
}

func NewHeaderFromFile(filePath string) (Header, error) {
	header := Header{}

	f, err := os.Open(filePath)
	if err != nil {
		return header, err
	}
	defer f.Close()

	// read given file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		// stop scanning on empty line
		if text == "" {
			break
		}

		err := header.ParseLine(text)
		if err != nil {
			return header, err
		}
	}

	if err := scanner.Err(); err != nil {
		return header, err
	}

	return header, nil
}

func NewHeaderFromConfig(config *Config) *Header {
	email, name := config.Archive.Email, config.Archive.Owner

	return &Header{
		ID:      uuid.New(),
		Date:    time.Now(),
		Authors: []Author{{Name: name, Email: email}},
	}
}
