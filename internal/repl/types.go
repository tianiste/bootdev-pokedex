package repl

import "pokedex/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Location []location
}

type response struct {
	Results           []location
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type location struct {
	Name string
	Url  string
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string
	Url  string
}
