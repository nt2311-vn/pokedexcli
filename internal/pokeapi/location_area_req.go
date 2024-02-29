package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasResp{}, fmt.Errorf("Bad status code: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
