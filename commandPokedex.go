package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("Pokedex empty")
	}
	fmt.Println("Pokedex: ")
	for key, _ := range cfg.caughtPokemon {
		fmt.Printf("  %s\n", key)
	}

	return nil
}
