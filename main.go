package main

import (
	"time"

	"github.com/stasdashkevitch/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapi: pokeClient,
	}
	startRepl(cfg)
}
