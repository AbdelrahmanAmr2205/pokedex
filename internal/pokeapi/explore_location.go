package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ExploreLocation(location string) ([]string, error) {
	url := baseURL + "/location-area/" + location
	var pokemons []string

	data, exists := c.cache.Get(url)
	if exists {
		pokemons = strings.Split(string(data), "\n")
	}

	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []string{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return []string{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return []string{}, err
		}

		locationResp := RespLocationArea{}
		err = json.Unmarshal(data, &locationResp)
		if err != nil {
			return []string{}, err
		}
		encounters := locationResp.PokemonEncounters
		pokemons = []string{}
		for _, encounter := range encounters {
			pokemons = append(pokemons, encounter.Pokemon.Name)
		}
		names := strings.Join(pokemons, "\n")
		c.cache.Add(url, []byte(names))
	}

	return pokemons, nil
}
