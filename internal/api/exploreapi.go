package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "pokedex-go/internal/cache"
)

type ExploreLocationResult struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string		`json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`	
}		



var exploreCache = pokecache.NewCache(5 * time.Second)

func FetchExploreResult(name string) {
	if data,ok := exploreCache.Get(name); ok {
		var cached ExploreLocationResult
		if err := json.Unmarshal(data, &cached); err == nil {
			for _, enc := range cached.PokemonEncounters {
				fmt.Println(enc.Pokemon.Name)
			}
			return
		}
	}
	
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", name)
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

	exploreCache.Add(name, body)

	var data ExploreLocationResult
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	for _, encounter := range data.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}