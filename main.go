package main

import (
	"time"

	"github.com/stasdashkevitch/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute)
	cfg := &config{
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		pokeapi:       pokeClient,
	}
	startRepl(cfg)
}
