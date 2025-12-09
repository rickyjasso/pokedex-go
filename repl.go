package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rickyjasso/pokedex-go/internal/pokeapi"
)

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command := input[0]

		if val, ok := commands[command]; !ok {
			fmt.Printf("Unknown command\n")
		} else {
			val.callback(conf)
		}
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)

	return result
}
