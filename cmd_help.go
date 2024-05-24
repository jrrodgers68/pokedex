package main

import "fmt"

func commandHelp(config *Config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	cmds := getCommands()
	for key := range cmds {
		fmt.Println(cmds[key].name, ": ", cmds[key].description)
	}
	return nil
}
