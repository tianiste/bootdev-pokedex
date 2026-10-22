package main

import (
	"pokedex/internal/pokecache"
	"pokedex/internal/repl"
	"time"
)

func main() {
	firstURL := "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	cache := pokecache.NewCache(5 * time.Second)

	cfg := repl.Config{
		Next:  &firstURL,
		Cache: cache,
	}

	repl.StartRepl(&cfg)
}
