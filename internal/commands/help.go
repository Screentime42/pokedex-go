package commands

import "fmt"

func help(args []string) {
	fmt.Println(`Welcome to the Pokedex!
Usage:

- help: Displays a help message
- exit: Exit the Pokedex
- map: Show next 20 location areas
- mapb: Show previous 20 location areas
- explore <location-area-name>: displays Pokemon encounters found in the area
- catch <pokemon-name>: attempt to catch a pokemon!
	`)
}