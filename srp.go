package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Single Responsibility Principle

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(_ int) {
	// ...
}

// separation of concerns

func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}

func SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}
