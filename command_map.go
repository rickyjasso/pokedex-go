package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Pokemap struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(c *config) error {
	if c.Next == nil {
		c.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	url := fmt.Sprintf("%v", c.Next)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	var pkmap Pokemap
	err = json.Unmarshal(body, &pkmap)
	if err != nil {
		fmt.Println(err)
	}

	c.Next = pkmap.Next
	c.Previous = pkmap.Previous

	for _, loc := range pkmap.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	url := fmt.Sprintf("%v", c.Previous)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error getting request: %v", err)
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil
	}
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	var pkmap Pokemap
	err = json.Unmarshal(body, &pkmap)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	c.Next = pkmap.Next
	c.Previous = pkmap.Previous

	for _, loc := range pkmap.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}
