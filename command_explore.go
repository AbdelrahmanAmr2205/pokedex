package main

import "fmt"

func commandExplore(conf *config, words []string) error {
	if len(words) == 1 {
		return fmt.Errorf("error: too few arguments\nUsage: %s {location}", words[0])
	} else if len(words) > 2 {
		return fmt.Errorf("error: too many arguments\nUsage: %s {location}", words[0])
	}

	pokemons, err := conf.pokeapiClient.ExploreLocation(words[1])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", words[1])
	if len(pokemons) == 0 {
		fmt.Println("No pokemons found")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Println(" -", pokemon)
	}

	return nil
}
