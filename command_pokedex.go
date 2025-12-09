package main

import "fmt"

func commandPokedex(c *config, arg string) error {

	fmt.Println("Your Pokedex:")
	for _, pkm := range c.Pokedex {
		fmt.Printf("\t- %s\n", pkm.Name)
	}
	return nil
}
