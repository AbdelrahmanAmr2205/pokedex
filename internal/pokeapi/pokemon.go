package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespPokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return RespPokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return RespPokemon{}, err
		}
	}

	pokemon := RespPokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return RespPokemon{}, err
	}

	return pokemon, nil
}
