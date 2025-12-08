package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
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
			fmt.Printf("Unknown command")
		} else {
			val.callback()
		}
	}

}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)

	return result
}
