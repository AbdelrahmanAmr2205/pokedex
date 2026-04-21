package main

import "fmt"

func commandHelp(conf *config, words []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n\n", command.name, command.description)
	}
	return nil
}
