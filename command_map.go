package main

import (
	"errors"
	"fmt"
)

func commandMapf(conf *config) error {
	if conf.nextLocationsURL == nil && conf.previousLocationsURL != nil {
		return errors.New("you're on the last page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.previousLocationsURL = locationsResp.Previous

	for _, locArea := range locationsResp.Results {
		fmt.Println(locArea.Name)
	}
	fmt.Println()

	return nil
}

func commandMapB(conf *config) error {
	if conf.nextLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.previousLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.previousLocationsURL = locationsResp.Previous

	for _, locArea := range locationsResp.Results {
		fmt.Println(locArea.Name)
	}
	fmt.Println()

	return nil
}
