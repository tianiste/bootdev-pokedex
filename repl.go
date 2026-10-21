package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() == false {
			return
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println("Error :", err)
		}
	}

}
