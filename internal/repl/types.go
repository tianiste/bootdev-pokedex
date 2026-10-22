package repl

import "pokedex/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

type response struct {
	Results []location
}

type location struct {
	Name string
	Url  string
}
