package pokeapi

import "encoding/json"

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName

	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)

	if ok {
		pokemon := Pokemon{}

		err := json.Unmarshal(data, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil

	}
}
