package main

import (
	"time"

	"github.com/rickyjasso/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	conf := &config{
		Client: pokeClient,
	}
	startRepl(conf)
}
