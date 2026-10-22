package repl

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func commandExit(cfg *Config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")
	commands := GetCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *Config, _ ...string) error {
	var resp response

	url := *cfg.Next
	body, exists := cfg.Cache.Get(url)
	if !exists {

		res, err := http.Get(*cfg.Next)
		if err != nil {
			fmt.Println("error while making get request", err)
			return err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, body)
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		cfg.Cache.Add(url, body)
	}
	err := json.Unmarshal(body, cfg)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	cfg.Location = append(cfg.Location, resp.Results...)

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *Config, _ ...string) error {
	var resp response

	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.Previous
	body, exists := cfg.Cache.Get(url)
	if !exists {

		res, err := http.Get(*cfg.Previous)
		if err != nil {
			fmt.Println("error while making get request", err)
			return err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, body)
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		cfg.Cache.Add(url, body)
	}
	err := json.Unmarshal(body, cfg)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	cfg.Location = append(cfg.Location, resp.Results...)
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location-name>")
	}
	name := args[0]
	var resp response

	//	found := slices.ContainsFunc(cfg.Location, func(loc location) bool {
	//		return loc.Name == name
	//	})

	//	if !found {
	//		fmt.Println("Location is not in the displayed list")
	//		return nil
	//	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", name)

	body, exists := cfg.Cache.Get(url)
	fmt.Printf("Exploring %s... \n", name)
	if !exists {

		res, err := http.Get(url)
		if err != nil {
			fmt.Println("error while making get request", err)
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, body)
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		cfg.Cache.Add(url, body)
	}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}

	return nil
}
func commandCatch(cfg *Config, args ...string) error {
	pokedex := cfg.Pokedex
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon-name>")
	}
	name := args[0]

	var resp pokemon

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error while making get request", err)
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)

	chance := rand.Intn(resp.BaseExperience) + 1
	if chance >= (resp.BaseExperience / 3) {
		pokedex[name] = resp
		fmt.Printf("%s was caught\nYou may now inspect it with the inspect command.\n", resp.Name)
		return nil
	}
	fmt.Printf("%s escaped!\n", resp.Name)

	return nil
}

func inspectCommand(cfg *Config, args ...string) error {
	pokedex := cfg.Pokedex
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}
	name := args[0]
	p, exists := pokedex[name]
	if !exists {
		return fmt.Errorf("you haven't caught %s", name)
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", p.Name, p.Height, p.Weight)
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}

func pokedexCommand(cfg *Config, args ...string) error {
	pokedex := cfg.Pokedex
	fmt.Println("Your pokedex:")
	for key := range pokedex {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			Callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the inputed location listed with the map command",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throws a ball at the pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays details of a caught Pokemon",
			Callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the caught Pokemon",
			Callback:    pokedexCommand,
		},
	}
}
