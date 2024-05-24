package main

import (
	"fmt"

	"gitlab.smarthomecomputing.net/john/pokedexcli/internal/pokeapi"
)

func printMapLocations(locations []pokeapi.Location) {

	// print out the list of locations in results
	for i := range locations {
		fmt.Println(locations[i].Name)
	}
}

func commandMap(config *Config, args []string) error {
	locs, err := config.Proxy.GetNextLocations()
	if err != nil {
		return nil
	}

	printMapLocations(locs)
	return nil
}

func commandMapB(config *Config, args []string) error {
	locs, err := config.Proxy.GetPreviousLocations()
	if err != nil {
		return nil
	}

	printMapLocations(locs)
	return nil
}
