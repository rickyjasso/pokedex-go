package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	return nil

}
