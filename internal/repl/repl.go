package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func StartRepl(cfg *Config) {
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

		cmd, ok := GetCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.Callback(cfg); err != nil {
			fmt.Println("Error :", err)
		}
	}

}
