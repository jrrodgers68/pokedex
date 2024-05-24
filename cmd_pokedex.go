package main

import "fmt"

func commandPokedex(config *Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for name := range config.Pokedex {
		fmt.Println(" - " + name)
	}
	return nil
}
