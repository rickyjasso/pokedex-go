package main

import (
	"fmt"
)

func commandMap(c *config) error {
	locations, err := c.Client.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locations, err := c.Client.ListLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locations.Next
	c.Previous = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
