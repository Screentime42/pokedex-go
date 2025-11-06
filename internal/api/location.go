package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count			int					`json:"count"`
	Next			string				`json:"next"`
	Previous 	*string				`json:"previous"`
	Results		[]LocationResult	`json:"results"`
}

type LocationResult struct {
	Name			string				`json:"name"`
	URL			string				`json:"url"`
}





func FetchLocationAreas(offset int) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=20&offset=%d", offset)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	var data LocationAreasResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	for _, loc := range data.Results {
		fmt.Printf("%s\n", loc.Name)
	}
}