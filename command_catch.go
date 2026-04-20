package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(conf *config, words []string) error {
	if len(words) == 1 {
		return fmt.Errorf("error: too few arguments\nUsage: %s {location}", words[0])
	} else if len(words) > 2 {
		return fmt.Errorf("error: too many arguments\nUsage: %s {location}", words[0])
	}

	pokemon, err := conf.pokeapiClient.GetPokemon(words[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", words[1])
	catchAttempt := rand.IntN(pokemon.BaseExperience)

	if catchAttempt <= 40 {
		fmt.Println(words[1], "was caught!")
		pokedex[pokemon.Name] = Pokemon{
			Name:       pokemon.Name,
			Experience: pokemon.BaseExperience,
		}
	} else {
		fmt.Println(words[1], "escaped!")
	}

	return nil
}
