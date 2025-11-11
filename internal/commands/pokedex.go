package commands

import (
	"fmt"
)

func commandPokedex(args []string) {
	if len(Pokedex) < 1 {
		fmt.Println("You havn't caught anything yet!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, e := range(Pokedex) {
			fmt.Printf("	- %s\n", e.Name)
		}
	}
}