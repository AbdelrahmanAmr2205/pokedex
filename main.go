package main

import (
	"time"

	"github.com/AbdelrahmanAmr2205/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	conf := &config{
		pokeapiClient: pokeapiClient,
	}
	startRepl(conf)
}
