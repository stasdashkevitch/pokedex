package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provided a location name")
	}

	name := args[0]
	location, err := cfg.pokeapi.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
