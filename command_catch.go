package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *config, name string) error {

	if name == "" {
		fmt.Println("Missing pokemon name: 'catch <pokemon>'")
		return nil
	}

	pokemon, err := c.Client.GetPokemon(name)
	if err != nil {
		fmt.Printf("%s is not a valid pokemon...\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	rn := rand.Float64()
	if rn > 0.8 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		_, exists := c.Pokedex[pokemon.Name]
		if !exists {
			c.Pokedex[name] = pokemon
			fmt.Printf("%s was added to your pokedex\n", name)

		} else {
			fmt.Printf("%s is already in your pokedex\n", name)
		}
	}

	return nil
}
