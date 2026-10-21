package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	Next     *string
	Previous *string
}

type response struct {
	Results []location
}

type location struct {
	Name string
	Url  string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	var resp response

	res, err := http.Get(*cfg.Next)
	if err != nil {
		fmt.Println("error while making get request", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed withs status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(body, cfg)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	var resp response
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(*cfg.Previous)
	if err != nil {
		fmt.Println("error while making get request", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed withs status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(body, cfg)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays pokemon locations",
			callback:    commandMapb,
		},
	}
}
