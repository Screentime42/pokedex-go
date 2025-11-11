package api


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	pokecache "pokedex-go/internal/cache"
)

type PokemonData struct {
	Name			string			`json:"name"`
	BaseXP		int				`json:"base_experience"`
	Height		int				`json:"height"`
	Weight		int				`json:"weight"`
	Stats			[]StatWrapper	`json:"stats"`
	Types			[]TypeWrapper	`json:"types"`
}


type StatWrapper struct {
	BaseStat 	int				`json:"base_stat"`
	Stat			StatDetails		`json:"stat"`
}

type StatDetails struct {
	Name 			string			`json:"name"`
}

type TypeWrapper struct {
	Type			TypeDetails		`json:"type"`
}

type TypeDetails struct {
	Name			string			`json:"name"`
}


var pokemonCache = pokecache.NewCache(1 * time.Hour)

func FetchPokemonResult(name string) (PokemonData, error) {
	
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", name)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return PokemonData{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return PokemonData{}, err
	}

	pokemonCache.Add(name, body)

	var data PokemonData
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return PokemonData{}, err
	}

	return data, nil
}
