package commands

import (
	"pokedex-go/internal/api"
	"fmt"
)

var page int = 0
const pageSize = 20

func locationAreas(args []string) {
	page++
	api.FetchLocationAreas((page - 1) * pageSize)
	
}

func prevLocationAreas(args []string) {
	if page > 1 {
		page--
		api.FetchLocationAreas((page - 1) * pageSize)
	} else {
		fmt.Println("You're on the first page.")
	}
}