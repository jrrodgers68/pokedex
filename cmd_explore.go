package main

import "fmt"

func commandExplore(config *Config, args []string) error {

	fmt.Println("Exploring " + args[0] + "...")

	la, err := config.Proxy.GetLocationArea(args[0])
	if err != nil {
		return nil
	}

	fmt.Println("Found Pokemon:")

	for i := 0; i < len(la.PokemonEncounters); i++ {
		name := la.PokemonEncounters[i].Pokemon.Name
		fmt.Println(" - " + name)
	}

	return nil
}
