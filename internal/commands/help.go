package commands

import "fmt"

func help() {
	fmt.Println(`Welcome to the Pokedex!
Usage:

- help: Displays a help message
- exit: Exit the Pokedex
- map: Show next 20 location areas
- mapb: Show previous 20 location areas`)
}