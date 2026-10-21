package main

import (
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func main() {
	firstURL := "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"

	cfg := *&config{
		Next: &firstURL,
	}
	startRepl(&cfg)
}
