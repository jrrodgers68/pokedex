package main

import (
	"fmt"

	"gitlab.smarthomecomputing.net/john/pokedexcli/internal/pokeapi"
)

func getPokemonStatByName(pokemon *pokeapi.Pokemon, statName string) int {
	for i := 0; i < len(pokemon.Stats); i++ {
		stat := pokemon.Stats[i]
		if stat.Stat.Name == statName {
			return pokemon.Stats[i].BaseStat
		}
	}

	return -1
}

func commandInspect(config *Config, args []string) error {
	name := args[0]

	p, ok := config.Pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Println("Name: " + p.Name)
		fmt.Println(fmt.Sprintf("Height: %d", p.Height))
		fmt.Println(fmt.Sprintf("Weight: %d", p.Weight))
		fmt.Println("Stats:")
		fmt.Println(fmt.Sprintf("   -hp: %d", getPokemonStatByName(&p, "hp")))
		fmt.Println(fmt.Sprintf("   -attack: %d", getPokemonStatByName(&p, "attack")))
		fmt.Println(fmt.Sprintf("   -defense: %d", getPokemonStatByName(&p, "defense")))
		fmt.Println(fmt.Sprintf("   -special-attack: %d", getPokemonStatByName(&p, "special-attack")))
		fmt.Println(fmt.Sprintf("   -special-defense: %d", getPokemonStatByName(&p, "special-defense")))
		fmt.Println(fmt.Sprintf("   -speed: %d", getPokemonStatByName(&p, "speed")))

		fmt.Println("Types:")
		for i := 0; i < len(p.Types); i++ {
			fmt.Println("  - " + p.Types[i].Type.Name)
		}

		return nil
	}
}
