package main

import "fmt"

func commandInspect(conf *config, words []string) error {
	if len(words) == 1 {
		return fmt.Errorf("error: too few arguments\nUsage: %s {location}", words[0])
	} else if len(words) > 2 {
		return fmt.Errorf("error: too many arguments\nUsage: %s {location}", words[0])
	}

	pokemon, err := conf.pokeapiClient.GetPokemon(words[1])
	if err != nil {
		return err
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println("  -", t.Type.Name)
	}

	return nil
}
