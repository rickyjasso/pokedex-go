package pokeapi

import (
	"encoding/json"
	"fmt"
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

	if resp.StatusCode == 404 {
		fmt.Println("Area doesn't exist...")
		return Pokemap{}, err
	}

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

func (c *Client) GetArea(areaToExplore string) (PokeArea, error) {
	url := baseURL + "/location-area/" + areaToExplore
	if val, ok := c.cache.Get(url); ok {
		area := PokeArea{}
		err := json.Unmarshal(val, &area)
		if err != nil {
			return PokeArea{}, err
		}
		return area, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeArea{}, err
	}

	area := PokeArea{}
	err = json.Unmarshal(data, &area)
	if err != nil {
		return PokeArea{}, err
	}
	c.cache.Add(url, data)
	return area, nil
}
