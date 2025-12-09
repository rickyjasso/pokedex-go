package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Pokemap, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locations := Pokemap{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return Pokemap{}, err
		}
		return locations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemap{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemap{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemap{}, err
	}

	locations := Pokemap{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Pokemap{}, err
	}

	c.cache.Add(url, data)
	return locations, nil

}
