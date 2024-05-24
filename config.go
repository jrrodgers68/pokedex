package main

import "gitlab.smarthomecomputing.net/john/pokedexcli/internal/pokeapi"

type Config struct {
	Proxy   *pokeapi.PokeAPIProxy
	Pokedex map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	p, _ := pokeapi.NewPokeAPIProxy()

	config := Config{Proxy: p, Pokedex: make(map[string]pokeapi.Pokemon, 0)}
	return &config
}

func (c *Config) StorePokemon(pokemon *pokeapi.Pokemon) error {
	c.Pokedex[pokemon.Name] = *pokemon
	return nil
}
