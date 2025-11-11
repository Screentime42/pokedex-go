package commands

import (
	"pokedex-go/internal/api"
	"fmt"
	"strings"
	"math/rand"
)

var Pokedex = make(map[string]api.PokemonData)


func commandCatch(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: catch <pokemon-name>")
		return
	}
	name := strings.TrimSpace(args[0])
	name = strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

	catch, err := api.FetchPokemonResult(name)
	if err != nil {
		fmt.Println("Failed to fetch Pokemon:", err)
		return
	}

	
	fmt.Printf("Throwing a Pokeball at %s...\n", name)



	if checkCatch(catch.BaseXP) {
		fmt.Printf("%s was caught!\n", name)
		key := strings.ToLower(name)
		Pokedex[key] = catch
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
}

func checkCatch(baseXP int) bool {
	catchProbability := 1.0 / (1.0 + float64(baseXP)/100.0)
	return rand.Float64() < catchProbability
}