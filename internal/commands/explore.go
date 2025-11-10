package commands

import (
	"pokedex-go/internal/api"
	"fmt"
	"strings"
)

func commandExplore(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: explore <location-area-name>, use map command to scroll through available areas")
	}
	name := strings.TrimSpace(args[0])

	api.FetchExploreResult(name)
}