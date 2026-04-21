package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AbdelrahmanAmr2205/pokedex/internal/pokeapi"
)

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command, ok := getCommands()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := command.callback(conf, words)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations(20 locations)",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations(20 locations)",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore  <location_name>",
			description: "Get the details of a certain location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Shows a pokemon's details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all caught pokemon",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string //url of the next page of resources
	previousLocationsURL *string //url of the previous page of resources
	pokedex              map[string]Pokemon
}

type Pokemon struct {
	Name       string
	Experience int
}
