package main

import (
	"fmt"
	"os"
)

const baseUrl string = "https://pokeapi.co/api/v2/"

func commandExit(conf *config, words []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, words []string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n\n", command.name, command.description)
	}
	return nil
}
