package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location in more detail",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a specific Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a Pokemon that you have already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect all caught Pokemon that you have stored in your Pokedex",
			callback:    commandPokedex,
		},
	}
}

func invokeCommand(cmd string, args []string, config *Config) error {
	cmds := getCommands()
	c, ok := cmds[cmd]
	if ok {
		return c.callback(config, args)
	} else {
		if len(cmd) == 0 {
			return nil
		}
		// invalid command - so print the help command
		fmt.Println("Command [", cmd, "] not found")
		return commandHelp(config, args)
	}
}
