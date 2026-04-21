package main

import "fmt"

func commandPokedex(conf *config, words []string) error {
	if len(words) > 1 {
		return fmt.Errorf("error: too many arguments\nUsage: %s {location}", words[0])
	}

	fmt.Println("Your Pokedex:")
	for key := range conf.pokedex {
		fmt.Println(" -", key)
	}

	return nil
}
