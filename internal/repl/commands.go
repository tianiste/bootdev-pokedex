package repl

import (
	"encoding/json"
	"fmt"
	"io"
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

	for _, location := range resp.Results {
		fmt.Println(location.Name)
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
	}
}
