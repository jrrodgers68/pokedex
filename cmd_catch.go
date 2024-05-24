package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args []string) error {
	if len(args) == 0 {
		fmt.Println("you must provide a pokemon name to catch")
		return nil
	}

	p, err := config.Proxy.GetPokemon(args[0])
	if err != nil {
		fmt.Println("unable to get Pokemon data!")
		return nil
	}

	fmt.Println("Throwing a Pokeball at " + p.Name + "...")

	// we got our pokemon --
	res := rand.Intn(p.BaseExperience)
	if res > 40 {
		fmt.Printf("%s escaped!\n", p.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", p.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	config.StorePokemon(p)

	return nil

}
