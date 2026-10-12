package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() == false {
			return
		}
		userInput := scanner.Text()
		cleanUserInput := cleanInput(userInput)
		if len(cleanUserInput) < 1 {
			continue
		}
		fmt.Printf("Your command was: %s\n", cleanUserInput[0])
	}
}
