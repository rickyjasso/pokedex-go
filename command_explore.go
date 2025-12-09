package main

import "fmt"

func commandExplore(c *config, areaToExplore string) error {
	area, err := c.Client.GetArea(areaToExplore)
	if err != nil {
		return err
	}
	if areaToExplore == "" {
		fmt.Printf("Missing area: 'explore <area>'\n")
		return nil
	}

	fmt.Printf("Exploring %s...\n", areaToExplore)
	fmt.Println("Found Pokemon:")

	for _, pkm := range area.PokemonEncounters {
		fmt.Printf("\t- %s\n", pkm.Pokemon.Name)
	}
	return nil
}
